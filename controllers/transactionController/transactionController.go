package transactioncontroller

import (
	"net/http"

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

	result := services.DepositService(c, req.Amount)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.NewErrorResponse(result.Error.Error()))
		return
	}

	response := transactionresponse.DepositResponse{
		Message: "Deposit successful",
		Success: true,
		Data: transactionresponse.CurrentBalance{
			AccountName:    result.Data.AccountName,
			AccountNumber:  result.Data.AccountNumber,
			CurrentBalance: result.Data.CurrentBalance,
		},
	}

	c.JSON(http.StatusOK, response)
}
