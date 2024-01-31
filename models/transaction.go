package models

import (
	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	ToAccount   uuid.UUID `json:"to_account"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Status      string    `json:"status"`

	User User `gorm:"foreignKey:UserID"`
}
