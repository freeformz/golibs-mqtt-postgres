package db

import (
	"fmt"

	"github.com/fltmtoda/golibs/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type sqlxDB struct {
	Dbm *sqlx.DB
}

type sqlxTxn struct {
	sqlxDB
	Tx *sqlx.Tx
}

func createSqlx(setting *Setting) (DB, error) {
	db, err := sqlx.Connect(
		setting.Dialect(),
		setting.URL,
	)
	if err != nil {
		return nil, err
	}
	log.Info("Open Database !")
	if setting.MaxIdleConns > 0 {
		db.DB.SetMaxIdleConns(setting.MaxIdleConns)
	}
	if setting.MaxOpenConns > 0 {
		db.DB.SetMaxOpenConns(setting.MaxOpenConns)
	}
	return &sqlxDB{
		Dbm: db,
	}, nil
}

/**************************
  gormDB functions
**************************/
func (db *sqlxDB) Create(value interface{}) error {
	// TODO
	return nil
}
func (db *sqlxDB) Save(value interface{}) error {
	// TODO
	return nil

}
func (db *sqlxDB) Delete(value interface{}, where ...interface{}) error {
	// TODO
	return nil
}

func (db *sqlxDB) AutoMigrate(values ...interface{}) error {
	// TODO
	return nil
}
func (db *sqlxDB) DropTableWithCascade(values ...interface{}) error {
	// TODO
	return nil
}

func (db *sqlxDB) CloseDB() error {
	if db == nil || db.Dbm == nil {
		return fmt.Errorf("Already closed db")
	}
	if err := db.Dbm.Close(); err != nil {
		log.Warn("Failed to close Database !")
	} else {
		log.Info("Close Database !")
	}
	return nil
}

func (db *sqlxDB) TxRunnable(f TxFunc) error {
	var err error
	txn, err := db.beginTransaction()
	if err != nil {
		return err
	}
	if err := f(txn); err != nil {
		txn.Rollback()
		return err
	}
	if err = txn.Commit(); err != nil {
		return err
	}
	return nil
}
func (db *sqlxDB) beginTransaction() (Transaction, error) {
	log.Debug("sqlx#Begin")
	var err error
	tx, err := db.Dbm.Beginx()
	if err != nil {
		log.Error("Failed to begin-transaction %v", err)
		return nil, err
	}
	txn := &sqlxTxn{}
	txn.Dbm = db.Dbm
	txn.Tx = tx
	return txn, nil
}
func (db *sqlxDB) Exec(sql string, values ...interface{}) error {
	ret, err := db.Dbm.Exec(sql, values...)
	_, err = ret.RowsAffected()
	return err
}

/**************************
  gormTxn functions
**************************/
func (txn *sqlxTxn) Exec(sql string, values ...interface{}) error {
	ret, err := txn.Tx.Exec(sql, values...)
	_, err = ret.RowsAffected()
	return err
}

func (txn *sqlxTxn) Commit() error {
	if txn.Tx == nil {
		return nil
	}
	log.Debug("sqlx#Commit")
	if err := txn.Tx.Commit(); err != nil {
		log.Error("Failed to commit-transaction %v", err)
		return err
	}
	txn.Dbm = nil
	return nil
}

func (txn *sqlxTxn) Rollback() error {
	if txn.Tx == nil {
		return nil
	}
	log.Warn("sqlx#Rollback")
	if err := txn.Tx.Rollback(); err != nil {
		log.Error("Failed to rollback-transaction %v", err)
		return err
	}
	txn.Dbm = nil
	return nil
}
