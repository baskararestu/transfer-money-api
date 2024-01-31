package routes

import (
	authcontroller "github.com/baskararestu/transfer-money/controllers/authController"
	"github.com/baskararestu/transfer-money/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/login", authcontroller.Login)
		authGroup.POST("/register", authcontroller.CreateUser)
		authGroup.POST("/logout", middlewares.AuthMiddleware(), authcontroller.Logout)
	}
}
