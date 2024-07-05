// Database connection and operations
// Author: Yuya KAMATAKI
// LICENSE: MIT

package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"os"   // for environment variables
	"time" // for using date
)

func ConnectDB() *gorm.DB {
	// Get DB credentials from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// Connect to the database
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=postgres port=" + dbPort + " sslmode=disable TimeZone=Asia/Tokyo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
