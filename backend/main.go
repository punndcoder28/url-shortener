package main

import "net/http"

func main() {
	http.HandleFunc("/shorten", func(writer http.ResponseWriter, request *http.Request) {
	})

	http.HandleFunc("/redirect", func(writer http.ResponseWriter, request *http.Request) {
	})
}

func handleShorten(writer http.ResponseWriter, request *http.Request) {
}

func handleRedirect(writer http.ResponseWriter, request *http.Request) {
}
