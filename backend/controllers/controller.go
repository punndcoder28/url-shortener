package controllers

import (
	"encoding/json"
	"fmt"

	helpers "github.com/punndcoder28/url-shortner/helpers"
	models "github.com/punndcoder28/url-shortner/models"
)

var dsn string

/*
CreateTables to create all the tables in the database
*/
func CreateTables() {
	db := helpers.ConnectDB()

	fmt.Println("Connection Opened to Database")

	helpers.CreateURLTable(db)

	fmt.Println("Database Migrated")
}

func insertURL(url string) {
	db := helpers.ConnectDB()

	hashedURL := helpers.HashURL(url)
	urlRecord := models.URL{Hash: hashedURL, URL: url}
	db.Create(&urlRecord)
}

func getURLFromHash(hash string) []byte {
	db := helpers.ConnectDB()

	var url models.URL
	db.Where("hash = ?", hash).First(&url)

	jsonURL, err := json.Marshal(url)
	if err != nil {
		panic(err)
	}

	return jsonURL
}
