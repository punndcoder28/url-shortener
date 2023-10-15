/*
Package helpers provides helper functions for the application.
*/
package helpers

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"github.com/punndcoder28/url-shortner/models"
)

/*
HandleErr - This function panics if there is an error
*/
func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*
PgError - This function returns a pg error
*/
func PgError(err error) *pgconn.PgError {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		return pgErr
	}

	return nil
}

/*
FindURLHashFromHash - This function finds the URL from the hash
*/
func FindURLHashFromHash(db *gorm.DB, hashString string) (string, error) {
	var url models.URL

	result := db.Where("hash = ?", hashString).First(&url)

	if result.Error != nil {
		return "", result.Error
	}

	urlString := url.Hash
	return urlString, nil
}
