package config

import (
	"log"

	migrate "github.com/Saadyyyy/money-bank-be/etc/models"
	"gorm.io/gorm"
)

func DBMigration(db *gorm.DB) {
	// Migrate User
	err := db.AutoMigrate(&migrate.Users{})
	if err != nil {
		log.Fatalf("Failed to migrate Category: %v", err)
	}

}
