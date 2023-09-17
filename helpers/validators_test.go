package helpers

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	models "github.com/punndcoder28/url-shortner/models"
)

func TestValidateShortenURLRequest(t *testing.T) {
	validRequest := models.CreateURLRequest{URL: "https://example.com"}

	err := ValidateShortenURLRequest(validRequest)
	assert.NoError(t, err, "Expected a valid request, but got an error")
}

func TestInvalidEmptyShortenURLRequest(t *testing.T) {
	invalidRequest := models.CreateURLRequest{URL: ""}
	err := ValidateShortenURLRequest(invalidRequest)
	assert.Error(t, err, "Expected an error for an invalid request, but got nil")

	validationErrors, ok := err.(validator.ValidationErrors)
	assert.True(
		t,
		ok,
		"Expected a validator.ValidationErrors type error, but got a different error type",
	)

	for _, e := range validationErrors {
		switch e.Field() {
		case "URL":
			assert.Equal(t, "required", e.Tag(), "Expected a required tag, but got a different tag")
		default:
			t.Errorf("Unknown validation error on field %s with tag %s", e.Field(), e.Tag())
		}
	}
}
