package database

import (
	"log"
	"os"
	"runtime"

	"github.com/carp-cobain/gin-todos/database/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectAndMigrate connects to a database and runs migrations using project models.
func ConnectAndMigrate() (*gorm.DB, error) {
	dsn := lookupConnectParams()
	db, err := Connect(dsn)
	if err != nil {
		return nil, err
	}
	if err = RunMigrations(db); err != nil {
		return nil, err
	}
	return db, nil
}

// Connect to a database. Either sqlite3 or postgres are supported.
func Connect(dsn string) (*gorm.DB, error) {
	config := &gorm.Config{
		Logger: logger.Discard, // disable gorm logger
	}
	db, err := gorm.Open(sqlite.Open(dsn), config)
	if err != nil {
		return nil, err
	}
	if err = optimize(db); err != nil {
		log.Printf("unable to optimize sqlite conn: %+v", err)
	}
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxOpenConns(max(4, runtime.NumCPU()))
	}
	return db, nil
}

// Run migrations on a database using project models.
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&model.Story{}, &model.Task{})
}

// Optimize a sqlite database for production.
func optimize(db *gorm.DB) error {
	return db.Exec(`PRAGMA journal_mode = WAL;
		PRAGMA busy_timeout = 5000;
		PRAGMA synchronous = NORMAL;
		PRAGMA cache_size = 1000000000;
		PRAGMA foreign_keys = true;
		PRAGMA temp_store = memory;`).Error
}

// Lookup db connection params from env vars
func lookupConnectParams() (dsn string) {
	dsn = "todos.db"
	if envar, ok := os.LookupEnv("DB_DSN"); ok {
		dsn = envar
	}
	return
}
