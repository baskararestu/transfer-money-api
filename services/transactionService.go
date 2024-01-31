package services

import (
	"errors"
	"fmt"

	"github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/middlewares"
	"github.com/baskararestu/transfer-money/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepositResponse struct {
	AccountName   string  `json:"account_name"`
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"`
}

type DepositServiceResult struct {
	DepositResponse
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

	if err := tx.Commit().Error; err != nil {
		return DepositServiceResult{Error: fmt.Errorf("failed to commit transaction: %v", err)}
	}

	response := DepositResponse{
		AccountName:   user.FullName,
		AccountNumber: bankAccount.AccountNumber,
		Balance:       bankAccount.Balance,
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
