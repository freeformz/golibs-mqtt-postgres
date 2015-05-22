package main

import (
	"fmt"
	"github.com/fltmtoda/golibs/db"
	"github.com/fltmtoda/golibs/db/sql"
	"github.com/fltmtoda/golibs/logger"
	"github.com/fltmtoda/golibs/mqtt"
	"github.com/tylerb/graceful"
	"net/http"
	"os"
	"strconv"
	"time"
)

var log = logger.GetLogger()

func main() {
	log.Info("Startup main process")

	// Setup mqtt client
	client := mqtt.Create(
		&mqtt.Setting{
			ClientId: fmt.Sprintf("mqtt-%v", time.Now().Unix()),
			URL:      os.Getenv("MQTT_BROKER_URL"),
		},
	)
	if err := client.Connect(); err != nil {
		log.Error(err)
		return
	}

	// Setup db
	maxIdleConns, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_IDLE_CONNS"))
	maxOpenConns, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_OPEN_CONNS"))
	dbm, err := db.Create(
		&db.Setting{
			Url:          os.Getenv("DATABASE_URL"),
			MaxIdleConns: maxIdleConns,
			MaxOpenConns: maxOpenConns,
		},
	)
	if err != nil {
		log.Error(err)
		return
	}

	// Setup db schema
	dbm.DropTableWithCascade(&MqttLog{})
	dbm.AutoMigrate(&MqttLog{})

	alive := true
	exit_server := make(chan int)
	exit_process := make(chan int)

	log.Info("Start http server...")
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		},
	)
	server := &graceful.Server{
		Timeout: 30 * time.Second,
		Server: &http.Server{
			Addr:    ":" + os.Getenv("PORT"),
			Handler: nil,
		},
	}
	server.ShutdownInitiated = func() {
		log.Info("Call ShutdownInitialated.")
		alive = false
		code := <-exit_server
		if code != 0 {
			log.Error("Faild to subscriber")
		}
		if client != nil {
			client.Disconnect(10)
			if dbm != nil {
				client = nil
			}
			dbm.CloseDB()
			dbm = nil
		}
		exit_process <- 0
	}
	go func() {
		server.ListenAndServe()
		log.Info("Stop http server")
	}()

	batchInsertCount, _ := strconv.Atoi(os.Getenv("PG_MAX_BATCH_INSERT_COUNT"))
	if batchInsertCount == 0 {
		batchInsertCount = 1
	}
	insertSQL := "INSERT INTO mqtt_log(message,created_at,updated_at)VALUES($1,$2,$3)"
	msgs := make([]string, 0, batchInsertCount)
	for {
		if !alive {
			exit_server <- 0
			break
		}
		err := client.Subscribe(
			"test/#",
			mqtt.QOS_ZERO,
			func(msg mqtt.Message) error {
				msgs = append(msgs, string(msg.Payload()))
				return nil
			},
		)
		if err != nil {
			log.Error(err)
			break
		}
		if batchInsertCount <= len(msgs) {
			go func(messages []string) {
				now := time.Now()
				err = dbm.TxRunnable(
					func(txn db.Transaction) error {
						for _, msg := range messages {
							log.Infof("Recieved message: %v", msg)
							err := txn.Exec(
								insertSQL,
								msg,
								now,
								now,
							)
							if err != nil {
								return err
							}
						}
						return nil
					},
				)
				if err != nil {
					log.Error(err)
				}
			}(msgs)
			msgs = msgs[0:0]
		}
	}
	code := <-exit_process
	log.Info("Shutdown main process")
	os.Exit(code)
}

type MqttLog struct {
	db.PK
	Message sql.StringType `json:"msg" sql:"type:json"`
	db.Timestamps
}
