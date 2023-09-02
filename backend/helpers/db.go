package helpers

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	models "github.com/punndcoder28/url-shortner/models"
)

func loadEnv() string {
	err := godotenv.Load()
	HandleErr(err)

	dsn := os.Getenv("DATABASE_DSN")

	return dsn
}

func ConnectDB() *gorm.DB {
	dsn := loadEnv()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	HandleErr(err)

	return db
}

func CreateURLTable(db *gorm.DB) {
	db.Migrator().DropTable(&models.URL{})
	db.AutoMigrate(&models.URL{})
}

func GetURLFromHash(hash string) (string, error) {
	db := ConnectDB()

	var url models.URL
	result := db.Where("hash = ?", hash).First(&url)

	if result.Error != nil {
		return "", result.Error
	}

	urlString := url.URL
	return urlString, nil
}
