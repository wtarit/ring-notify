package main

import (
	"api/configs"
	"api/models"
)

func main() {
	configs.InitDatabase()
	db := configs.DB()
	db.AutoMigrate(&models.User{})
}
