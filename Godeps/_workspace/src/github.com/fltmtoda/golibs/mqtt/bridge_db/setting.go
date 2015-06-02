package bdb

import (
	"github.com/fltmtoda/golibs/db"
	"github.com/fltmtoda/golibs/mqtt"
)

type Setting struct {
	Topic       string
	Concurrency int
	DB          *db.Setting
	MQTT        *mqtt.Setting
}
