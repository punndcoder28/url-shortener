package helpers

import (
	"encoding/hex"
	"errors"
	"hash/fnv"

	models "github.com/punndcoder28/url-shortner/models"
)

/*
HashURL - This function returns a hashed string of the URL
*/
func HashURL(url string) string {
	hash := fnv.New32a()
	hash.Write([]byte(url))
	hashValue := hash.Sum(nil)

	hashHex := hex.EncodeToString(hashValue)

	return hashHex
}

/*
CreateTempURLFromHash - This function returns a temporary URL from the hash
*/
func CreateTempURLFromHash(hash string) string {
	return "http://localhost:8080/" + hash
}

/*
InsertURL - This function inserts a URL into the database
*/
func InsertURL(url string) (string, error) {
	db := ConnectDB()

	hashedURL := HashURL(url)
	urlRecord := models.URL{Hash: hashedURL, URL: url}
	result := db.Create(&urlRecord)
	if result.Error != nil {
		if pgErr := PgError(result.Error); pgErr != nil {
			if pgErr.Code == "23505" {
				url, err := FindURLHashFromHash(db, hashedURL)
				if err != nil {
					return "", errors.New(result.Error.Error())
				}

				return url, nil
			}
		}

		return "", errors.New(result.Error.Error())
	}

	return hashedURL, nil
}
