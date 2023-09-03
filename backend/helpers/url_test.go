package helpers

import (
	"testing"
)

func TestHashURL(t *testing.T) {
	url := "https://example.com"
	hash := HashURL(url)
	expectedHash := "6fbc04d3"
	if hash == "" {
		t.Errorf("HashURL(%s) return empty hash", url)
	}

	if hash != expectedHash {
		t.Errorf("HashURL(%s) = %s; want %s", url, hash, "d5c3dbf6")
	}
}

func TestCreateTempURLFromHash(t *testing.T) {
	hash := "d5c3dbf6"
	expectedURL := "http://localhost:8080/d5c3dbf6"
	tempURL := CreateTempURLFromHash(hash)

	if tempURL != expectedURL {
		t.Errorf("CreateTempURLFromHash(%s) = %s; want %s", hash, tempURL, expectedURL)
	}
}

func TestInsertURL(t *testing.T) {
	validURL := "https://example.com"
	hash1, err := InsertURL(validURL)
	if err != nil {
		t.Errorf("InsertURL(%s) returned an error: %v", validURL, err)
	}

	if hash1 == "" {
		t.Errorf("InsertURL(%s) returned an empty hash", validURL)
	}

	duplicateURL := "https://example.com"
	hash2, err := InsertURL(duplicateURL)
	if err == nil && hash2 == "" {
		t.Errorf("InsertURL(%s) inserted duplicate URL: %v", duplicateURL, err)
	}
}
