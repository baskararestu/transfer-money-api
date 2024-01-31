package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string

const (
	DepositTransaction  TransactionType = "deposit"
	TransferTransaction TransactionType = "transfer"
	PaymentTransaction  TransactionType = "payment"
)

type TransactionHistory struct {
	ID              uuid.UUID       `gorm:"type:uuid;primaryKey" json:"id"`
	UserID          uuid.UUID       `json:"user_id"`
	User            User            `gorm:"foreignKey:UserID"`
	TransactionType TransactionType `json:"transaction_type"`
	Amount          float64         `json:"amount"`
	Description     string          `json:"description"`
	Timestamp       time.Time       `json:"timestamp"`
}
