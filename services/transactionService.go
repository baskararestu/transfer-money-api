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

// DepositResponse represents the response structure for the deposit service.
type DepositResponse struct {
	Balance     float64 `json:"balance"`
	AccountName string  `json:"account_name"`
}

// DepositServiceResult represents the result of the deposit service.
type DepositServiceResult struct {
	DepositResponse
	Error error
}

// DepositService handles the logic for depositing money into a user's bank account.
func DepositService(c *gin.Context, amount float64) DepositServiceResult {
	// Validate JWT token and extract user ID
	userID, err := middlewares.GetUserIDFromContext(c)
	if err != nil {
		return DepositServiceResult{Error: err}
	}

	// Start a database transaction
	tx := database.DB.Begin()

	// Defer the transaction rollback if there's an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check if user exists
	var user models.User
	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		return DepositServiceResult{Error: errors.New("user not found")}
	}

	// Deposit money into user's bank account
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

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return DepositServiceResult{Error: fmt.Errorf("failed to commit transaction: %v", err)}
	}

	// Construct the response
	response := DepositResponse{
		Balance:     bankAccount.Balance,
		AccountName: user.FullName,
	}

	return DepositServiceResult{DepositResponse: response}
}

// getBankAccountByUserID retrieves the bank account associated with the given user ID.
func getBankAccountByUserID(tx *gorm.DB, userID uuid.UUID) (*models.BankAccount, error) {
	var bankAccount models.BankAccount
	if err := tx.Where("user_id = ?", userID).First(&bankAccount).Error; err != nil {
		return nil, err
	}
	return &bankAccount, nil
}
