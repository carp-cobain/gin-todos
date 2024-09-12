package database

import (
	"os"

	"github.com/carp-cobain/gin-todos/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

// ConnectAndMigrate connects to a db and runs migrations using project models.
// Both sqlite3 and postgres are supported.
func ConnectAndMigrate() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	dialect, dsn := getConnectParams()
	config := &gorm.Config{
		Logger: logger.Discard, // disable gorm logger
	}
	if dialect == "postgres" {
		db, err = gorm.Open(postgres.Open(dsn), config)
	} else {
		db, err = gorm.Open(sqlite.Open(dsn), config)
	}
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&model.Story{}, &model.Task{}); err != nil {
		return nil, err
	}
	return db, nil
}
