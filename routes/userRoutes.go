package routes

import (
	authcontroller "github.com/baskararestu/transfer-money/controllers/authController"
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
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/login", authcontroller.Login)
		authGroup.POST("/register", authcontroller.CreateUser)
	}
}
