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

func Transfer(c *gin.Context) {
	var req struct {
		ReceiverAccountNumber string  `json:"receiver_account_number"`
		Amount                float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorresponse.NewErrorResponse(err.Error()))
		return
	}

	result := services.TransferService(c, req.ReceiverAccountNumber, req.Amount)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, errorresponse.NewErrorResponse(result.Error.Error()))
		return
	}

	transferData := transactionresponse.TransferData{
		Amount:       req.Amount,
		FromAccount:  result.Data.FromAccount,
		SenderName:   result.Data.SenderName,
		ToAccount:    req.ReceiverAccountNumber,
		ReceiverName: result.Data.ReceiverName,
	}

	response := transactionresponse.TransferResponse{
		Message: "Transfer successful",
		Success: true,
		Data:    transferData,
	}

	c.JSON(http.StatusOK, response)
}
