package main

import (
	"fmt"

	"github.com/punndcoder28/url-shortner/controllers"
)

func main() {
	fmt.Println("Welcome to URL Shortner")

	fmt.Println("Creating tables")
	controllers.CreateTables()

	fmt.Println("Created tables")
}
