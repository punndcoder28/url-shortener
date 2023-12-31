package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashURL(t *testing.T) {
	url := "https://example.com"
	hash := HashURL(url)
	expectedHash := "6fbc04d3"

	assert.NotEqual(t, "", hash, "HashURL returned empty hash")

	assert.Equal(t, expectedHash, hash, "HashURL returned correct hash")
}

func TestCreateTempURLFromHash(t *testing.T) {
	hash := "d5c3dbf6"
	expectedURL := "http://localhost:8080/d5c3dbf6"
	tempURL := CreateTempURLFromHash(hash)

	assert.Equal(t, expectedURL, tempURL, "CreateTempURLFromHash returned correct temp URL")
}

func TestInsertURL(t *testing.T) {
	validURL := "https://example.com"
	hash, err := InsertURL(validURL)

	assert.NoError(t, err, "No error while inserting URL")
	assert.NotEqual(t, "", hash, "Hash is empty for string")
}

func TestInsertDuplicateURL(t *testing.T) {
	duplicateURL := "https://example.com"
	hash, err := InsertURL(duplicateURL)

	assert.Error(t, err, "Expected error while inserting duplicate URL")
	assert.Equal(t, "", hash, "Expected empty hash for duplicate URL")
}
