package models

import (
	"github.com/google/uuid"
)

type TransactionType string

const (
	DepositTransaction  TransactionType = "deposit"
	TransferTransaction TransactionType = "transfer"
	PaymentTransaction  TransactionType = "payment"
)

type Transaction struct {
	ID          uuid.UUID       `gorm:"type:uuid;primaryKey" json:"id"`
	AccountID   uuid.UUID       `json:"account_id"`
	ToAccountID uuid.UUID       `json:"to_account_id"`
	Amount      float64         `json:"amount"`
	Type        TransactionType `json:"type"`
	Status      string          `json:"status"`

	Account BankAccount `gorm:"foreignKey:AccountID"`

	ToAccount BankAccount `gorm:"foreignKey:ToAccountID"`
}
