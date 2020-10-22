package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB global variable
var DB *gorm.DB

// OpenConnection - open sqlite connection
func OpenConnection() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	database, err := gorm.Open(sqlite.Open("./db/companies.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Category{},
		&CategoryBusiness{},
		&Customer{},
		&Company{},
		&Market{},
		&Order{},
		&OrderItem{},
		&Person{},
		&Product{},
		&Role{},
		&ShoppingCart{},
		&User{})
	DB = database
}
