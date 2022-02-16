package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func Example_listRepositories() {
	dsn := os.Getenv("MYSQL_DSN")
	if len(dsn) == 0 {
		panic("MYSQL_DSN is empty")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	results, err := listRepositories(db)
	if err != nil {
		panic(err)
	}
	_, _ = fmt.Fprintln(os.Stderr, len(results))
	// Output:
}
