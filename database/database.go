package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/baskararestu/transfer-money/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "postgres://postgres:postgres@localhost:5432/db_transfer_money"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := database.AutoMigrate(&models.User{}, &models.BankAccount{}, &models.Transaction{}); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	DB = database

	log.Println("Connected to database!")
}
