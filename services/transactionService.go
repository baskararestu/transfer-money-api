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

func TransferService(c *gin.Context, receiverAccountNumber string, amount float64) transactionresponse.TransferServiceResult {
	userID, err := middlewares.GetUserIDFromContext(c)
	if err != nil {
		return transactionresponse.TransferServiceResult{Error: err}
	}

	tx := database.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	senderAccount, err := getBankAccountByUserID(tx, userID)
	if err != nil {
		tx.Rollback()
		return transactionresponse.TransferServiceResult{Error: fmt.Errorf("failed to get sender's bank account: %v", err)}
	}

	senderUser, err := GetUserByID(userID.String())
	if err != nil {
		tx.Rollback()
		return transactionresponse.TransferServiceResult{Error: fmt.Errorf("failed to get sender's user data: %v", err)}
	}
	senderName := senderUser.FullName

	receiverAccount, err := getBankAccountByAccountNumber(tx, receiverAccountNumber)
	if err != nil {
		tx.Rollback()
		return transactionresponse.TransferServiceResult{Error: fmt.Errorf("failed to get receiver's bank account: %v", err)}
	}

	if senderAccount.Balance < amount {
		tx.Rollback()
		return transactionresponse.TransferServiceResult{Error: errors.New("insufficient balance")}
	}

	senderAccount.Balance -= amount
	if err := tx.Save(&senderAccount).Error; err != nil {
		tx.Rollback()
		return transactionresponse.TransferServiceResult{Error: fmt.Errorf("failed to update sender's bank account balance: %v", err)}
	}

	receiverAccount.Balance += amount
	if err := tx.Save(&receiverAccount).Error; err != nil {
		tx.Rollback()
		return transactionresponse.TransferServiceResult{Error: fmt.Errorf("failed to update receiver's bank account balance: %v", err)}
	}

	// Retrieve receiver's name
	receiverName := ""
	if err := tx.Model(&receiverAccount.User).Select("full_name").Scan(&receiverName).Error; err != nil {
		tx.Rollback()
		return transactionresponse.TransferServiceResult{Error: fmt.Errorf("failed to retrieve receiver's name: %v", err)}
	}

	if err := createTransferTransactions(tx, senderAccount, receiverAccount, amount); err != nil {
		tx.Rollback()
		return transactionresponse.TransferServiceResult{Error: fmt.Errorf("failed to create transaction records: %v", err)}
	}

	data := transactionresponse.TransferData{
		Amount:       amount,
		FromAccount:  senderAccount.AccountNumber,
		ToAccount:    receiverAccount.AccountNumber,
		ReceiverName: receiverName,
		SenderName:   senderName,
	}

	if err := tx.Commit().Error; err != nil {
		return transactionresponse.TransferServiceResult{Error: fmt.Errorf("failed to commit transaction: %v", err)}
	}

	return transactionresponse.TransferServiceResult{
		Success: true,
		Message: "Transfer successful",
		Data:    data,
	}
}

func getBankAccountByAccountNumber(tx *gorm.DB, accountNumber string) (*models.BankAccount, error) {
	var bankAccount models.BankAccount
	if err := tx.Where("account_number = ?", accountNumber).First(&bankAccount).Error; err != nil {
		return nil, err
	}
	return &bankAccount, nil
}

func createTransferTransactions(tx *gorm.DB, sender, receiver *models.BankAccount, amount float64) error {
	setTransaction := models.Transaction{
		ID:          uuid.New(),
		AccountID:   sender.ID,
		ToAccountID: receiver.ID,
		Amount:      amount,
		Type:        models.TransferTransaction,
		Status:      "completed",
	}

	if err := createTransaction(tx, &setTransaction); err != nil {
		return err
	}

	return nil
}
func createTransaction(tx *gorm.DB, transaction *models.Transaction) error {
	if err := tx.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}
