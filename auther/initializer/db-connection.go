package initializer

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB
func LoadDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err= gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal("Faild connect to DB")
	}
}