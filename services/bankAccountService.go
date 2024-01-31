// services/bank_account_service.go

package services

import (
	"time"

	"github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateBankAccountForUser(userID uuid.UUID) (uuid.UUID, error) {
	bankAccountID := uuid.New()

	bankAccount := models.BankAccount{
		ID:        bankAccountID,
		UserID:    userID,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := database.DB.Create(&bankAccount).Error; err != nil {
		return uuid.Nil, err
	}

	return bankAccountID, nil
}

func CreateBankAccountForUserInTransaction(tx *gorm.DB, userID string) (uuid.UUID, error) {
	bankAccountID := uuid.New()

	bankAccount := models.BankAccount{
		ID:        bankAccountID,
		UserID:    uuid.MustParse(userID),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := tx.Create(&bankAccount).Error; err != nil {
		return uuid.Nil, err
	}

	return bankAccountID, nil
}

func GetBankAccountByID(bankAccountID uuid.UUID) (*models.BankAccount, error) {
	var bankAccount models.BankAccount
	if err := database.DB.Where("id = ?", bankAccountID).First(&bankAccount).Error; err != nil {
		return nil, err
	}
	return &bankAccount, nil
}
