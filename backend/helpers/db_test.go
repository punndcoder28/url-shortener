package helpers

import (
	"testing"

	models "github.com/punndcoder28/url-shortner/models"
)

func TestConnectDB(t *testing.T) {
	db := ConnectDB()

	if db == nil {
		t.Errorf("ConnectDB returned nil, expected *gorm.DB")
	}
}

func TestCreateURLTable(t *testing.T) {
	db := ConnectDB()

	CreateURLTable(db)

	if !db.Migrator().HasTable(&models.URL{}) {
		t.Errorf("CreateURLTable did not create the URL table")
	}
}

func TestGetURLFromHash(t *testing.T) {
	db := ConnectDB()

	testURL := models.URL{Hash: "testHash", URL: "https://example.com"}
	db.Create(&testURL)

	retrievedURL, err := GetURLFromHash("testHash")
	if err != nil {
		t.Errorf("GetURLFromHash returned an error: %v", err)
	}

	expectedURL := "https://example.com"
	if retrievedURL != expectedURL {
		t.Errorf("GetURLFromHash returned %s, expected %s", retrievedURL, expectedURL)
	}
}
