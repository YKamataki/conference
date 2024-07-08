// Database connection and operations
// Author: Yuya KAMATAKI
// LICENSE: MIT

package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"
	"os"   // for environment variables
	"time" // for using date
)

func ConnectDB() *gorm.DB {
	// Get DB credentials from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Connect to the database
	dsn := fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

type Conference struct {
	ID          uint
	Title       string
	Description string
	Date        time.Time
	Passcode    *string
	CreatedAt   time.Time
	Presenters  []Presenter
}

type Presenter struct {
	ID           uint
	ConferenceID uint
	Name         string
	Topic        *string
	Order        uint
}

func MigrateDB(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&Conference{}, &Presenter{})
}

func GetConferences(db *gorm.DB) []Conference {
	// Get all GetConferences
	var conferences []Conference
	db.Preload("Presenters").Find(&conferences)
	return conferences
}
