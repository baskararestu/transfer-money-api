package transactioncontroller

import (
	"net/http"

	"github.com/baskararestu/transfer-money/middlewares"
	"github.com/baskararestu/transfer-money/services"
	"github.com/gin-gonic/gin"
)

// DepositResponse represents the response structure for the deposit transaction.
type DepositResponse struct {
	Message string         `json:"message"`
	Success bool           `json:"success"`
	Data    CurrentBalance `json:"data"`
}

type CurrentBalance struct {
	AccountName    string  `json:"account_name"`
	AccountNumber  string  `json:"account_number"`
	CurrentBalance float64 `json:"current_balance"`
}

func Deposit(c *gin.Context) {
	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := middlewares.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID from context"})
		return
	}

	result := services.DepositService(c, req.Amount)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if err := services.AddTransactionToHistory(userID, services.DepositTransaction, req.Amount, "Deposit"); err != nil {
		rollbackErr := services.RollbackDepositTransaction(userID, req.Amount)
		if rollbackErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add deposit transaction to history and rollback the deposit transaction"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add deposit transaction to history"})
		return
	}

	response := DepositResponse{
		Message: "Deposit successful",
		Success: true,
		Data: CurrentBalance{
			AccountName:    result.DepositResponse.AccountName,
			AccountNumber:  result.DepositResponse.AccountNumber,
			CurrentBalance: result.DepositResponse.Balance,
		},
	}

	c.JSON(http.StatusOK, response)
}
