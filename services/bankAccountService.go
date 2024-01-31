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
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomDigits := strconv.Itoa(randomGenerator.Intn(1000000000))

	randomDigits = fmt.Sprintf("%09v", randomDigits)

	return "101" + randomDigits
}

func RollbackDepositTransaction(userID uuid.UUID, amount float64) error {
	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	bankAccount, err := getBankAccountByUserID(tx, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	bankAccount.Balance -= amount

	if err := tx.Save(&bankAccount).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
