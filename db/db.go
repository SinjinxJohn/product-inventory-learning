package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func InitDB(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}
	return db, nil
}
