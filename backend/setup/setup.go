// Package setup : This package contains all the setup functions
package setup

import (
	"fmt"

	helpers "github.com/punndcoder28/url-shortner/helpers"
)

func CreateTables() {
	db := helpers.ConnectDB()

	fmt.Println("Connection Opened to Database")

	helpers.CreateURLTable(db)

	fmt.Println("Database Migrated")
}
