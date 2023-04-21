package database

import (
	"log"
	"os"

	"github.com/sazzo/sazztv/backend/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func CreateConnection() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.StreamSettings{})
	db.AutoMigrate(&model.StreamCredentials{})

	if err != nil {
		log.Fatal(err)
	}
}

func GetConnection() *gorm.DB {
	return db
}