package configs

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDatabase() {
	dsn := os.Getenv("DB_CONN_STRING")
	var err error
	database, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("error connecting to database: %v \n", err)
	}
}

func DB() *gorm.DB {
	return database
}
