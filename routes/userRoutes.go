package routes

import (
	authcontroller "github.com/baskararestu/transfer-money/controllers/authController"
	userController "github.com/baskararestu/transfer-money/controllers/userController"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/api/users")
	{
		userGroup.GET("/", userController.Index)
		userGroup.GET("/:id", userController.Show)
		userGroup.PUT("/:id", userController.Update)
		userGroup.DELETE("/:id", userController.Delete)
	}
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/login", authcontroller.Login)
		authGroup.POST("/register", authcontroller.CreateUser)
	}
}
