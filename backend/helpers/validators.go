// Package helpers : This package contains all the validator helper functions
package helpers

import (
	"github.com/go-playground/validator/v10"

	models "github.com/punndcoder28/url-shortner/models"
)

var validate *validator.Validate

// ValidateShortenURLRequest : This function validates the CreateURLRequest
func ValidateShortenURLRequest(createURLRequest models.CreateURLRequest) error {
	validate = validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(createURLRequest)
}
