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
func ConnectAndMigrate() (*gorm.DB, *gorm.DB, error) {
	dsn := dsnEnvLookup()
	writer, err := Connect(dsn, 1)
	if err != nil {
		return nil, nil, err
	}
	if err := RunMigrations(writer); err != nil {
		return nil, nil, err
	}
	reader, err := Connect(dsn, max(4, runtime.NumCPU()))
	if err != nil {
		return nil, nil, err
	}
	return reader, writer, nil
}

// Connect to a sqlite3 database.
func Connect(dsn string, maxConns int) (*gorm.DB, error) {
	config := &gorm.Config{
		Logger: logger.Discard, // disable gorm logger
	}
	db, err := gorm.Open(sqlite.Open(dsn), config)
	if err != nil {
		return nil, err
	}
	if err = setPragmas(db); err != nil {
		log.Printf("unable to set PRAGMAs for sqlite connection: %+v", err)
	}
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxOpenConns(maxConns)
	}
	return db, nil
}

// Run migrations on a database using project models.
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&model.Story{}, &model.Task{})
}

// Optimize a sqlite database for production.
func setPragmas(db *gorm.DB) error {
	return db.Exec(`PRAGMA journal_mode = WAL;
		PRAGMA busy_timeout = 5000;
		PRAGMA synchronous = NORMAL;
		PRAGMA cache_size = 1000000000;
		PRAGMA foreign_keys = true;
		PRAGMA temp_store = memory;`).Error
}

// Lookup db dsn param from env var
func dsnEnvLookup() string {
	dsn, ok := os.LookupEnv("DB_DSN")
	if !ok {
		log.Panicf("DB_DSN not defined")
	}
	return dsn
}
