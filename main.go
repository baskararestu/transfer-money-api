package main

import (
	"github.com/gin-gonic/gin"

	db "enigma.com/learn-golang/database"
	"enigma.com/learn-golang/routes"
)

func main() {
	port := "8080"

	r := gin.Default()
	db.ConnectDatabase()

	routes.ProductRoutes(r)
	routes.UserRoutes(r)

	r.GET("/api-1", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})
	r.Run(":" + port)
}
