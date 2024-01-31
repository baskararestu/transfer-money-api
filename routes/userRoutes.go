package routes

import (
	userController "github.com/baskararestu/transfer-money/controllers/userController"
	"github.com/baskararestu/transfer-money/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/api/users")
	userGroup.Use(middlewares.AuthMiddleware())
	{
		userGroup.GET("/", userController.Index)
		userGroup.GET("/:id", userController.Show)
		userGroup.PUT("/:id", userController.Update)
	}
}
