package database

import (
	"log"
	"os"

	"github.com/autumnleaf-ra/go-blogger-app/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env")
	}

	dsn := os.Getenv("DSN")
	database, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	} else {
		log.Println("Connected!")
	}

	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
