package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is an abstraction over the database
var DB *gorm.DB

// Get db connection params from env vars
func getConnectParams() (dialect string, dsn string) {
	dialect, dsn = "sqlite3", "todos.db"
	if envar, ok := os.LookupEnv("DB_DIALECT"); ok {
		dialect = envar
	}
	if envar, ok := os.LookupEnv("DB_DSN"); ok {
		dsn = envar
	}
	return
}

// ConnectDbAndMigrate connects to a db and runs migrations using project models.
// Both sqlite3 and postgres are supported.
func ConnectDbAndMigrate() {
	var db *gorm.DB
	var err error
	config := &gorm.Config{
		Logger: logger.Discard, // disable gorm logger
	}

	dialect, dsn := getConnectParams()
	if dialect == "postgres" {
		db, err = gorm.Open(postgres.Open(dsn), config)
	} else {
		db, err = gorm.Open(sqlite.Open(dsn), config)
	}

	if err != nil {
		log.Panicf("Failed to open database: %+v", err)
	}

	db.AutoMigrate(&Story{})
	DB = db
}
