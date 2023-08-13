package controllers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	models "github.com/punndcoder28/url-shortner/models"
)

const dsn = "postgresql://puneeth:5RHMSaa5GqZdCjM7uFXVjQ@projects-prod-3242.7s5.cockroachlabs.cloud:26257/url_shortener?sslmode=verify-full"

/*
CreateTables to create all the tables in the database
*/
func CreateTables() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	createURLTable(db)

	fmt.Println("Database Migrated")
}

func createURLTable(db *gorm.DB) {
	db.AutoMigrate(&models.URL{})
}
