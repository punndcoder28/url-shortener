package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	helpers "github.com/punndcoder28/url-shortner/helpers"
	models "github.com/punndcoder28/url-shortner/models"
)

func insertURL(url string) string {
	db := helpers.ConnectDB()

	hashedURL := helpers.HashURL(url)
	fmt.Println("Hashed URL: ", hashedURL)
	urlRecord := models.URL{Hash: hashedURL, URL: url}
	db.Create(&urlRecord)

	return hashedURL
}

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
	_ = json.NewDecoder(r.Body).Decode(&createURLRequest)

	hash := insertURL(createURLRequest.URL)
	url := helpers.CreateTempURLFromHash(hash)
	response := map[string]interface{}{"url": url}
	json.NewEncoder(w).Encode(response)
}
