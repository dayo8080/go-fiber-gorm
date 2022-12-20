package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	var err error
	const POSTGRES = "host=localhost user=gofiber password=gofiber dbname=go_fiber_gorm port=5432 sslmode=disable TimeZone=Africa/Lagos"
	dsn := POSTGRES
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}
	DB = db
	fmt.Println("Connected to Database")
}
