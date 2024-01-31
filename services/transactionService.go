package services

import (
	"errors"
	"fmt"

	"github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/middlewares"
	"github.com/baskararestu/transfer-money/models"
	transactionresponse "github.com/baskararestu/transfer-money/responses/transactionResponse"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepositServiceResult struct {
	transactionresponse.DepositResponse
	Error error
}

func DepositService(c *gin.Context, amount float64) DepositServiceResult {
	userID, err := middlewares.GetUserIDFromContext(c)
	if err != nil {
		return DepositServiceResult{Error: err}
	}

	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user models.User
	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		return DepositServiceResult{Error: errors.New("user not found")}
	}

	bankAccount, err := getBankAccountByUserID(tx, userID)
	if err != nil {
		tx.Rollback()
		return DepositServiceResult{Error: fmt.Errorf("failed to get bank account: %v", err)}
	}

	bankAccount.Balance += amount
	if err := tx.Save(&bankAccount).Error; err != nil {
		tx.Rollback()
		return DepositServiceResult{Error: fmt.Errorf("failed to update bank account balance: %v", err)}
	}

	transaction := models.Transaction{
		ID:          uuid.New(),
		AccountID:   bankAccount.ID,
		ToAccountID: bankAccount.ID,
		Amount:      amount,
		Type:        models.DepositTransaction,
		Status:      "completed",
	}

	if err := createTransaction(tx, &transaction); err != nil {
		tx.Rollback()
		return DepositServiceResult{Error: fmt.Errorf("failed to create transaction: %v", err)}
	}

	if err := tx.Commit().Error; err != nil {
		return DepositServiceResult{Error: fmt.Errorf("failed to commit transaction: %v", err)}
	}

	response := transactionresponse.DepositResponse{
		Message: "Deposit successful",
		Success: true,
		Data: transactionresponse.CurrentBalance{
			AccountName:    user.FullName,
			AccountNumber:  bankAccount.AccountNumber,
			CurrentBalance: bankAccount.Balance,
		},
	}
	return DepositServiceResult{DepositResponse: response}
}

func getBankAccountByUserID(tx *gorm.DB, userID uuid.UUID) (*models.BankAccount, error) {
	var bankAccount models.BankAccount
	if err := tx.Where("user_id = ?", userID).First(&bankAccount).Error; err != nil {
		return nil, err
	}
	return &bankAccount, nil
}

func createTransaction(tx *gorm.DB, transaction *models.Transaction) error {
	if err := tx.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}
