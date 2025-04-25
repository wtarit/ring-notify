package main

import (
	"api/configs"
	"api/models"
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
	db.AutoMigrate(&models.User{})
}
