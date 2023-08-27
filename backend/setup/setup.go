package setup

import (
	"fmt"

	helpers "github.com/punndcoder28/url-shortner/helpers"
)

func main() {
	createTables()
}

func createTables() {
	db := helpers.ConnectDB()

	fmt.Println("Connection Opened to Database")

	helpers.CreateURLTable(db)

	fmt.Println("Database Migrated")
}
