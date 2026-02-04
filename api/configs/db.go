package configs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDatabase() {
	dsn := os.Getenv("DB_CONN_STRING")
	var err error
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{DryRun: false})
	if err != nil {
		log.Fatalf("error connecting to database: %v \n", err)
	}
	fmt.Println("Done connecting to DB.")
}

func DB() *gorm.DB {
	return database
}
