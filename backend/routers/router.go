package routers

import (
	"github.com/gorilla/mux"

	controllers "github.com/punndcoder28/url-shortner/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/shorten", controllers.CreateURL).Methods("POST")
	router.HandleFunc("/{shortUrl}", controllers.GetURL).Methods("GET")

	return router
}
