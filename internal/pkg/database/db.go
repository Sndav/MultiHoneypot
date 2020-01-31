package database

import (
	"Multi-Honeypot/internal/pkg/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

type DB struct {
	_gorm *gorm.DB
}

func NewDB(dbType string, url string) (*DB, error) {
	db := &DB{}
	var err error
	db._gorm, err = gorm.Open(dbType, url)
	return db, err
}

func (db *DB) InsertLog(logType string, err ...interface{}) error {
	var errs error
	logString := fmt.Sprint(err...)
	v := db._gorm.Create(&models.Logs{
		LogType: logType,
		Log:     logString,
	})
	errs = v.Error
	return errs
}

func (db *DB) InsertPhishing(pluginName string, fromIP string, payload string) error {
	var errs error
	v := db._gorm.Create(&models.Victim{
		FromIP:  fromIP,
		Service: pluginName,
		Payload: payload,
	})
	errs = v.Error
	return errs
}

func (dbs *DB) Migrator() error {
	db := dbs._gorm
	if v := db.AutoMigrate(&models.Victim{}); v.Error != nil {
		db.Rollback()
		return v.Error
	}
	if v := db.AutoMigrate(&models.Logs{}); v.Error != nil {
		db.Rollback()
		return v.Error
	}
	return nil
}
