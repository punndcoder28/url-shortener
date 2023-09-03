package helpers

import (
	"testing"

	"github.com/go-playground/validator/v10"

	models "github.com/punndcoder28/url-shortner/models"
)

func TestValidateShortenURLRequest(t *testing.T) {
	validRequest := models.CreateURLRequest{URL: "https://example.com"}
	invalidRequest := models.CreateURLRequest{URL: ""}

	err := ValidateShortenURLRequest(validRequest)
	if err != nil {
		t.Errorf("Expected a valid request, but got an error: %v", err)
	}

	err = ValidateShortenURLRequest(invalidRequest)
	if err == nil {
		t.Error("Expected an error for an invalid request, but got nil")
	} else {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			t.Error("Expected a validator.ValidationErrors type error, but got a different error type")
		} else {
			for _, e := range validationErrors {
				switch e.Field() {
				case "URL":
					switch e.Tag() {
					case "required":
						t.Skipf("URL is required, but it's missing")
					case "url":
						t.Errorf("URL is not a valid URL: %s", e.Value())
					default:
						t.Errorf("Validation error on field %s with tag %s", e.Field(), e.Tag())
					}
				default:
					t.Errorf("Unknown validation error on field %s with tag %s", e.Field(), e.Tag())
				}
			}
		}
	}
}
