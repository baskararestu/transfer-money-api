package main

import (
	productController "enigma.com/learn-golang/controllers/productController"
	"github.com/gin-gonic/gin"

	db "enigma.com/learn-golang/database"
)

func main() {
	r := gin.Default()
	db.ConnectDatabase()

	r.GET("/api/products", productController.Index)
	r.GET("/api/products/:id", productController.Show)
	r.POST("/api/products", productController.Create)
	r.PUT("/api/products/:id", productController.Update)
	r.DELETE("/api/products/:id", productController.Delete)

	r.Run()
}
