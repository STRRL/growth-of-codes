package repository

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	dsn := os.Getenv("MYSQL_DSN")
	if len(dsn) == 0 {
		fmt.Fprintf(w, "MYSQL_DSN is not set")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		_, _ = fmt.Fprintf(w, "gorm.Open error: %v", err)
		return
	}
	repositories, err := listRepositories(db)
	if err != nil {
		_, _ = fmt.Fprintf(w, "listRepositories error: %v", err)
		return
	}
	bytes, err := json.Marshal(repositories)
	if err != nil {
		_, _ = fmt.Fprintf(w, "json.Marshal error: %v", err)
		return
	}
	_, _ = w.Write(bytes)
}

func listRepositories(db *gorm.DB) ([]string, error) {
	var result []string
	query := db.Raw("select distinct project from file_complexity_snapshots;").Scan(&result)
	if query.Error != nil {
		return nil, query.Error
	}
	return result, nil
}
