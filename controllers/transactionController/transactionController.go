package transactioncontroller

import (
	"net/http"

	"github.com/baskararestu/transfer-money/services"
	"github.com/gin-gonic/gin"
)

// DepositResponse represents the response structure for the deposit transaction.
type DepositResponse struct {
	Message        string  `json:"message"`
	Success        bool    `json:"success"`
	CurrentBalance float64 `json:"current_balance"`
	AccountName    string  `json:"account_name"`
}

// Deposit handles the deposit transaction.
func Deposit(c *gin.Context) {
	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the deposit service to process the transaction
	result := services.DepositService(c, req.Amount)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Construct the response
	response := DepositResponse{
		Message:        "Deposit successful",
		Success:        true,
		CurrentBalance: result.DepositResponse.Balance,
		AccountName:    result.DepositResponse.AccountName,
	}

	c.JSON(http.StatusOK, response)
}
