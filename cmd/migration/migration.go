package config

import (
	"log"

	migrate "github.com/Saadyyyy/money-bank-be/etc/models"
	"gorm.io/gorm"
)

func DBMigration(db *gorm.DB) {
	// Migrate User
	err1 := db.AutoMigrate(&migrate.Users{})
	if err1 != nil {
		log.Fatalf("Failed to migrate Category: %v", err1)
	}

	//Migrate Card
	err2 := db.AutoMigrate(&migrate.Card{})
	if err2 != nil {
		log.Fatalf("Failed to migrate Category: %v", err2)
	}
}
