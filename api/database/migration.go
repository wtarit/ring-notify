package main

import (
	"api/configs"
	"api/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}

	configs.InitDatabase()
	db := configs.DB()

	fmt.Println("Starting database migration...")

	// Auto-migrate all models
	fmt.Println("Creating/updating tables...")
	if err := db.AutoMigrate(&models.Device{}, &models.APIKey{}); err != nil {
		log.Fatalf("error auto migrating database: %v\n", err)
	}

	fmt.Println("Migration completed successfully!")
	fmt.Println("Database schema is ready for use.")
}
