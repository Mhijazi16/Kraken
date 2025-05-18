package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load .env file, error: %s", err)
	}

	dsn := os.Getenv("dsn")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	var servers []Server
	if err := db.Find(&servers).Error; err != nil {
		log.Fatal("Query failed:", err)
	}
}
