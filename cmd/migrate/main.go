package main

import (
	"log"
	"os"

	"example.com/event-app/config"
	"example.com/event-app/db"

	mysqlConfig "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // needed for "file://" source
)

func main() {
	// initialize DB connection
	dbConn, err := db.InitDB(mysqlConfig.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		Addr:                 config.Envs.DBAddress,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}

	// create migration driver from DB connection
	driver, err := mysql.WithInstance(dbConn, &mysql.Config{})
	if err != nil {
		log.Fatal("Could not create migration driver:", err)
	}

	// run migrations
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal("Could not create migrate instance:", err)
	}

	cmd := os.Args[(len(os.Args) - 1)]

	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration failed:", err)
		} else {
			log.Println("Up Executed Successfully")
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration failed:", err)
		} else {
			log.Println("Down Executed Successfully")
		}
	}

}
