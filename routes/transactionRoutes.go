package routes

import (
	transactioncontroller "github.com/baskararestu/transfer-money/controllers/transactionController"
	"github.com/baskararestu/transfer-money/middlewares"
	"github.com/gin-gonic/gin"
)

// SetTransactionRoutes sets up the transaction routes.
func TransactionRoutes(router *gin.Engine) {
	transaction := router.Group("/api/transaction")
	transaction.Use(middlewares.AuthMiddleware())

	{
		transaction.POST("/deposit", transactioncontroller.Deposit)
		// Add more routes for other transaction operations if needed
	}
}
