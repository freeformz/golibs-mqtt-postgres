package db

import (
	"github.com/fltmtoda/golibs/logger"
)

var (
	log = logger.GetLogger()
)

type (
	DB interface {
		Create(value interface{}) error
		Save(value interface{}) error
		Delete(value interface{}, where ...interface{}) error

		AutoMigrate(values ...interface{}) error
		DropTableWithCascade(values ...interface{}) error

		TxRunnable(f TxFunc) error
		beginTransaction() (Transaction, error)

		Exec(sql string, values ...interface{}) error

		CloseDB() error
	}

	Transaction interface {
		Create(value interface{}) error
		Save(value interface{}) error
		Delete(value interface{}, where ...interface{}) error

		Exec(sql string, values ...interface{}) error

		Commit() error
		Rollback() error
	}

	Query interface {
	}

	TxFunc func(txn Transaction) error
)

func Create(setting *Setting) (DB, error) {
	return createGorm(setting)
}
