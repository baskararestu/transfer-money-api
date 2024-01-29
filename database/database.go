package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"enigma.com/learn-golang/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "postgres://postgres:postgres@localhost:5432/db_learn_golang"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := database.AutoMigrate(&models.Product{}, &models.User{}); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	DB = database

	log.Println("Connected to database!")
}
