package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDBPostgres initializes the PostgreSQL database connection using GORM.
func InitDBPostgres(cfg *AppConfig) (*gorm.DB, error) {
	// Create the DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHOST, cfg.DBUSERNAME, cfg.DBPASSWORD, cfg.DBNAME, cfg.DBPORT,
	)

	// Attempt to open a connection to the database
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Disable prepared statement caching for simple queries
	}), &gorm.Config{
		PrepareStmt: true, // Enable prepared statements for performance
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	log.Println("Database connected successfully")
	return db, nil
}
