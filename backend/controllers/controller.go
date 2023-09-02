// Package controllers : This package contains all the controllers
package controllers

import (
	"encoding/json"
	"net/http"

	helpers "github.com/punndcoder28/url-shortner/helpers"
	models "github.com/punndcoder28/url-shortner/models"
)

/*
GetURL - This function returns the URL from the hash
*/
func GetURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hash := r.URL.Path[1:]
	url, err := helpers.GetURLFromHash(hash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]interface{}{"error": "URL not found"}
		json.NewEncoder(w).Encode(response)
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}

/*
CreateURL - This function creates a URL
*/
func CreateURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var createURLRequest models.CreateURLRequest

	json.NewDecoder(r.Body).Decode(&createURLRequest)

	validationErr := helpers.ValidateShortenURLRequest(createURLRequest)
	if validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{"Message": "Invalid request body sent"}
		json.NewEncoder(w).Encode(response)
		return
	}

	hash := helpers.InsertURL(createURLRequest.URL)
	url := helpers.CreateTempURLFromHash(hash)
	response := map[string]interface{}{"url": url}
	json.NewEncoder(w).Encode(response)
}
