package main

import (
	"fmt"
	"log"

	migration "github.com/Saadyyyy/money-bank-be/cmd/migration"
	configs "github.com/Saadyyyy/money-bank-be/etc/config"
	"github.com/Saadyyyy/money-bank-be/pkg/routers"
	"github.com/go-redis/redis"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()

	// Initialize GORM DB
	gormDB, err := configs.InitDBPostgres(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.REDISADDRES, // e.g., "localhost:6379"
		Password: "",              // No password set
		DB:       0,               // Default DB
	})

	// Test Redis connection
	_, err = rdb.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Retrieve underlying *sql.DB from GORM
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("Failed to get *sql.DB instance from GORM: %v", err)
	}

	// Convert *sql.DB to *sqlx.DB
	sqlxDB := sqlx.NewDb(sqlDB, "postgres")

	// Initialize Echo instance
	e := echo.New()

	//migration table
	migration.DBMigration(gormDB)

	// Middleware setup
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Serve static files
	e.Static("/assets", "assets")

	// Register routes with Echo
	routers.Register(sqlxDB, e, rdb)

	port := fmt.Sprintf(":%d", cfg.SERVERPORT)
	log.Printf("Starting server on port %s ", port)
	e.Logger.Fatal(e.Start(port))
}
