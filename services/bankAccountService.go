// services/bank_account_service.go

package services

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateBankAccountForUser(tx *gorm.DB, userID string, accountNumber string) (uuid.UUID, error) {
	bankAccountID := uuid.New()

	bankAccount := models.BankAccount{
		ID:            bankAccountID,
		AccountNumber: accountNumber,
		UserID:        uuid.MustParse(userID),
		Balance:       0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
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

func GenerateAccountNumber() string {
	// Initialize a new local random generator with the current time as seed
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate random 9-digit number
	randomDigits := strconv.Itoa(randomGenerator.Intn(1000000000))

	// Pad randomDigits with leading zeros to ensure it's 9 digits long
	randomDigits = fmt.Sprintf("%09v", randomDigits)

	// Combine "101" with randomDigits to form a 12-digit account number
	return "101" + randomDigits
}
