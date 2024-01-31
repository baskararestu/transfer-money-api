package main

import (
	"github.com/gin-gonic/gin"

	db "github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/routes"
)

func main() {
	port := "5500"

	r := gin.Default()
	db.ConnectDatabase()

	routes.UserRoutes(r)

	r.GET("/api-1", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})
	r.Run(":" + port)
}
