package services

import (
	"time"

	"github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/models"
	"github.com/google/uuid"
)

type TransactionType string

const (
	DepositTransaction  TransactionType = "deposit"
	TransferTransaction TransactionType = "transfer"
	PaymentTransaction  TransactionType = "payment"
)

func AddTransactionToHistory(userID uuid.UUID, transactionType TransactionType, amount float64, description string) error {
	transaction := models.TransactionHistory{
		ID:              uuid.New(),
		UserID:          userID,
		TransactionType: models.TransactionType(transactionType),
		Amount:          amount,
		Description:     description,
		Timestamp:       time.Now(),
	}

	if err := database.DB.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}
