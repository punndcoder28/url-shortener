package main

import (
	"fmt"
	"log"
	"net/http"

	routers "github.com/punndcoder28/url-shortner/routers"
)

func main() {
	fmt.Println("Welcome to URL Shortner")

	fmt.Println("Creating tables")

	fmt.Println("Created tables")

	fmt.Println("Starting server")

	r := routers.Router()

	log.Fatal(http.ListenAndServe(":8080", r))
}
