package routes

import (
	authcontroller "enigma.com/learn-golang/controllers/authController"
	userController "enigma.com/learn-golang/controllers/userController"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/api/users")
	{
		userGroup.GET("/", userController.Index)
		userGroup.GET("/:id", userController.Show)
		userGroup.POST("/", userController.CreateUser)
		userGroup.PUT("/:id", userController.Update)
		userGroup.DELETE("/:id", userController.Delete)
	}
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/login", authcontroller.Login)
	}
}
