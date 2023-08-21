package main

import (
	"fmt"
	"log"
	"net/http"

	controllers "github.com/punndcoder28/url-shortner/controllers"
	routers "github.com/punndcoder28/url-shortner/routers"
)

func main() {
	fmt.Println("Welcome to URL Shortner")

	fmt.Println("Creating tables")
	controllers.CreateTables()

	fmt.Println("Created tables")

	fmt.Println("Starting server")

	r := routers.Router()

	log.Fatal(http.ListenAndServe(":8080", r))
}
