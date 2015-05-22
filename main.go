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
	"time"
)

var log = logger.GetLogger()

func main() {
	log.Info("Startup main process")

	// Setup mqtt client
	client := mqtt.Create(
		&mqtt.Setting{
			ClientId: fmt.Sprintf("mqtt-%v", time.Now().Unix()),
			URL:      os.Getenv("CLOUDMQTT_URL"),
		},
	)
	if err := client.Connect(); err != nil {
		log.Error(err)
		return
	}

	// Setup db
	dbm, err := db.Create(
		&db.Setting{
			Url:          os.Getenv("DATABASE_URL"),
			MaxIdleConns: 3,
			MaxOpenConns: 10,
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
		Timeout: 10 * time.Second,
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
			client.Disconnect(100)
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

	insertSQL := "INSERT INTO mqtt_log(message,created_at,updated_at)VALUES($1,$2,$3)"
	go func() {
		msgs := make([]string, 0, 10)
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
			if len(msgs) > 0 {
				go func(messages []string) {
					now := time.Now()
					err = dbm.TxRunnable(
						func(txn db.Transaction) error {
							for _, msg := range messages {
								log.Infof("Recieved message: %v", msg)
								err := dbm.Exec(
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
	}()
	code := <-exit_process
	log.Info("Shutdown main process")
	os.Exit(code)
}

type MqttLog struct {
	db.PK
	Message sql.StringType `json:"msg" sql:"type:json"`
	db.Timestamps
}
