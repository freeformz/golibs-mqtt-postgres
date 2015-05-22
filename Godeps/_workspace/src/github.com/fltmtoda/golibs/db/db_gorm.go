package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type gormDB struct {
	Dbm *gorm.DB
}

type gormTxn struct {
	gormDB
}

func createGorm(setting *Setting) (DB, error) {
	db, err := gorm.Open(
		setting.Dialect(),
		setting.Url,
	)
	if err != nil {
		return nil, err
	}
	log.Info("Open Database !")

	db.SingularTable(true)
	if log.IsDebugEnabled() {
		db.LogMode(true)
	}
	db.DB()
	if setting.IsDefaultSetting() {
		db.DB().SetMaxIdleConns(5)
		db.DB().SetMaxOpenConns(10)
	} else {
		db.DB().SetMaxIdleConns(setting.MaxIdleConns)
		db.DB().SetMaxOpenConns(setting.MaxOpenConns)
	}
	return &gormDB{
		Dbm: &db,
	}, nil
}

/**************************
  gormDB functions
**************************/
func (db *gormDB) Create(value interface{}) error {
	return db.Dbm.Create(value).Error
}
func (db *gormDB) Save(value interface{}) error {
	return db.Dbm.Save(value).Error
}
func (db *gormDB) Delete(value interface{}, where ...interface{}) error {
	return db.Dbm.Delete(value, where...).Error
}

func (db *gormDB) AutoMigrate(values ...interface{}) error {
	if len(values) > 0 {
		for _, v := range values {
			if err := db.Dbm.AutoMigrate(v).Error; err != nil {
				return err
			}
			log.Infof(
				"Create table: %s",
				db.Dbm.NewScope(v).TableName(),
			)
		}
	}
	return nil
}
func (db *gormDB) DropTableWithCascade(values ...interface{}) error {
	if len(values) > 0 {
		for _, v := range values {
			tableName := db.Dbm.NewScope(v).TableName()
			sql := fmt.Sprintf(
				"DROP TABLE IF EXISTS %s CASCADE",
				tableName,
			)
			if err := db.Dbm.Exec(sql).Error; err != nil {
				return err
			}
			log.Infof(
				"Drop table: %s",
				tableName,
			)
		}
	}
	return nil
}

func (db *gormDB) CloseDB() error {
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

func (db *gormDB) TxRunnable(f TxFunc) error {
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
func (db *gormDB) beginTransaction() (Transaction, error) {
	log.Debug("gorm#Begin")
	txn := &gormTxn{}
	txn.Dbm = db.Dbm.Begin()
	if err := txn.Dbm.Error; err != nil {
		log.Error("Failed to begin-transaction %v", err)
		return nil, err
	}
	return txn, nil
}
func (db *gormDB) Exec(sql string, values ...interface{}) error {
	return db.Dbm.Exec(sql, values...).Error
}

/**************************
  gormTxn functions
**************************/
func (txn *gormTxn) Commit() error {
	if txn.Dbm == nil {
		return nil
	}
	log.Debug("gorm#Commit")
	if err := txn.Dbm.Commit().Error; err != nil {
		log.Error("Failed to commit-transaction %v", err)
		return err
	}
	txn.Dbm = nil
	return nil
}

func (txn *gormTxn) Rollback() error {
	if txn.Dbm == nil {
		return nil
	}
	log.Warn("gorm#Rollback")
	if err := txn.Dbm.Rollback().Error; err != nil {
		log.Error("Failed to rollback-transaction %v", err)
		return err
	}
	txn.Dbm = nil
	return nil
}
