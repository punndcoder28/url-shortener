/*
Package routers contains all the routes for the application
*/
package routers

import (
	"github.com/gorilla/mux"

	controllers "github.com/punndcoder28/url-shortner/controllers"
)

/*
Router - This function returns a pointer to a mux.Router we can use as a handler.
*/
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/shorten", controllers.CreateURL).Methods("POST")
	router.HandleFunc("/{shortUrl}", controllers.GetURL).Methods("GET")

	return router
}
