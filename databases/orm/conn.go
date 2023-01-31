package orm

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", dbHost, dbUser, dbPass, dbName)

	DB, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return DB, nil
}
