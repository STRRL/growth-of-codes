package serverless

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func getGormDB() (*gorm.DB, error) {
	if db == nil {
		var err error
		db, err = initGormDBFromEnv()
		return db, err
	}
	return db, nil
}

func initGormDBFromEnv() (*gorm.DB, error) {
	// load dsn from OS env
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return nil, errors.New("env DB_DSN is not set")
	}
	result, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return result, err
}
