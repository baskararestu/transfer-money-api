package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	models "enigma.com/learn-golang/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/db_learn_golang"))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := database.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	DB = database

	log.Println("Connected to database!")
}
