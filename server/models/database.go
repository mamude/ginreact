package models

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func OpenConnection() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	database, err := gorm.Open(sqlite.Open("movies.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(&Company{}, &Person{}, &Role{}, &User{})
	DB = database
}
