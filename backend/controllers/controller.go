package controllers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	models "github.com/punndcoder28/url-shortner/models"
)

var dsn string

/*
CreateTables to create all the tables in the database
*/
func CreateTables() {
	loadEnv()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	createURLTable(db)

	fmt.Println("Database Migrated")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dsn = os.Getenv("DATABASE_DSN")
}

func createURLTable(db *gorm.DB) {
	db.AutoMigrate(&models.URL{})
}
