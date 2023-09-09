package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	models "github.com/punndcoder28/url-shortner/models"
)

func TestConnectDB(t *testing.T) {
	db := ConnectDB()

	assert.NotEqual(t, nil, db, "ConnectDB returned nil, expected *gorm.DB")
}

func TestCreateURLTable(t *testing.T) {
	db := ConnectDB()

	CreateURLTable(db)

	hasTable := db.Migrator().HasTable(&models.URL{})

	assert.Equal(t, true, hasTable, "Table creation failed! URL table not present")
}

func TestGetURLFromHash(t *testing.T) {
	db := ConnectDB()

	testURL := models.URL{Hash: "testHash", URL: "https://example.com"}
	db.Create(&testURL)

	retrievedURL, err := GetURLFromHash("testHash")
	assert.NoError(t, err, "No error while retrieving URL from hash")

	expectedURL := "https://example.com"
	assert.Equal(t, expectedURL, retrievedURL, "Correct URL retrieved from hash")
}
