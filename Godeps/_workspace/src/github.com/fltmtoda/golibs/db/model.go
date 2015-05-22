package db

import (
	"time"
	//	"github.com/fltmtoda/golibs/db/sql"
)

type (
	PK struct {
		Id uint64 `json:"-" gorm:"primary_key"`
	}
	Timestamps struct {
		CreatedAt time.Time `json:"-" sql:"type:timestamp without time zone;not null"`
		UpdatedAt time.Time `json:"-" sql:"type:timestamp without time zone;not null"`
	}
)
