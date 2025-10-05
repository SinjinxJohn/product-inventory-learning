package main

import (
	"database/sql"
	"log"

	"example.com/event-app/cmd/api"
	"example.com/event-app/config"
	"example.com/event-app/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.InitDB(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		Addr:                 config.Envs.DBAddress,
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	initStorage(db)

	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Failed to ping db:", err)
	}

	log.Println("Successfully connected to db")
}
