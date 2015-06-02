package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fltmtoda/golibs/db"
	"github.com/fltmtoda/golibs/db/sql"
	"github.com/fltmtoda/golibs/env"
	"github.com/fltmtoda/golibs/log"
	"github.com/fltmtoda/golibs/mqtt"
	"github.com/fltmtoda/golibs/mqtt/bridge_db"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/stretchr/graceful"
)

var (
	OK = map[string]bool{
		"ok":           true,
		"acknowledged": true,
	}
)

func main() {
	initSchema()

	setting := &bdb.Setting{
		Topic:       env.GetEnv("MQTT_TOPIC", "test/#").String(),
		Concurrency: int(env.GetEnv("MQTT_CONCURRENCY", "10").Int64()),
		DB: &db.Setting{
			URL:          os.Getenv("DATABASE_URL"),
			MaxIdleConns: int(env.GetEnv("DATABASE_MAX_IDLE_CONNS", "10").Int64()),
			MaxOpenConns: int(env.GetEnv("DATABASE_MAX_OPEN_CONNS", "10").Int64()),
		},
		MQTT: &mqtt.Setting{
			ClientId: fmt.Sprintf("Bridge-DB-%v", time.Now().Unix()),
			URL:      env.GetEnv("MQTT_BROKER_URL", "tcp://localhost:1883").String(),
		},
	}
	bridgeDB, err := bdb.Create(setting)
	if err != nil {
		panic(err)
	}
	go bridgeDB.Start(
		func(topic string) interface{} {
			return &RawData{}
		},
	)

	api := rest.NewApi()
	api.Use(rest.DefaultCommonStack...)
	router, err := rest.MakeRouter(
		&rest.Route{
			HttpMethod: "GET",
			PathExp:    "/",
			Func: func(w rest.ResponseWriter, r *rest.Request) {
				w.WriteJson(OK)
			},
		},
	)
	if err != nil {
		panic(err)
	}
	api.SetApp(router)

	server := &graceful.Server{
		Server: &http.Server{
			Addr:    ":" + env.GetEnv("PORT", "9000").String(),
			Handler: api.MakeHandler(),
		},
		Timeout:          10 * time.Second,
		NoSignalHandling: false,
		ShutdownInitiated: func() {
			log.Info("Call ShutdownInitiated")
			bridgeDB.Stop()
		},
	}
	server.ListenAndServe()
}

func initSchema() {
	dbm, err := db.Create(
		&db.Setting{
			URL: os.Getenv("DATABASE_URL"),
		},
	)
	if err != nil {
		panic(err)
	}
	defer dbm.CloseDB()
	dbm.DropTableWithCascade(&RawData{})
	dbm.AutoMigrate(&RawData{})
}

type RawData struct {
	db.PK
	DeviceId  sql.StringType    `json:"device_id"  sql:"not null"`
	Body      sql.StringType    `json:"body"       sql:"not null;type:json"`
	Timestamp sql.TimestampType `json:"created_at" sql:"type:timestamp without time zone"`
	db.Timestamps
}
