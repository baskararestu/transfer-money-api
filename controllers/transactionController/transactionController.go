package transactioncontroller

import (
	"net/http"

	"github.com/baskararestu/transfer-money/middlewares"
	errorresponse "github.com/baskararestu/transfer-money/responses/errorResponse"
	transactionresponse "github.com/baskararestu/transfer-money/responses/transactionResponse"
	"github.com/baskararestu/transfer-money/services"
	"github.com/gin-gonic/gin"
)

func Deposit(c *gin.Context) {
	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.NewErrorResponse(err.Error()))
		return
	}

	userID, err := middlewares.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.NewErrorResponse("Failed to get user ID from context"))
		return
	}

	result := services.DepositService(c, req.Amount)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.NewErrorResponse(result.Error.Error()))
		return
	}

	if err := services.AddTransactionToHistory(userID, services.DepositTransaction, req.Amount, "Deposit"); err != nil {
		rollbackErr := services.RollbackDepositTransaction(userID, req.Amount)
		if rollbackErr != nil {
			c.JSON(http.StatusInternalServerError, errorresponse.NewErrorResponse("Failed to add deposit transaction to history and rollback the deposit transaction"))
			return
		}

		c.JSON(http.StatusInternalServerError, errorresponse.NewErrorResponse("Failed to add deposit transaction to history"))
		return
	}

	response := transactionresponse.DepositResponse{
		Message: "Deposit successful",
		Success: true,
		Data: transactionresponse.CurrentBalance{
			AccountName:    result.DepositResponse.AccountName,
			AccountNumber:  result.DepositResponse.AccountNumber,
			CurrentBalance: result.DepositResponse.Balance,
		},
	}

	c.JSON(http.StatusOK, response)
}
