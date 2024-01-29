// router.go
package routes

import (
	productController "enigma.com/learn-golang/controllers/productControllers"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/api/products", productController.Index)
	r.GET("/api/products/:id", productController.Show)
	r.POST("/api/products", productController.Create)
	r.PUT("/api/products/:id", productController.Update)
	r.DELETE("/api/products/:id", productController.Delete)
}
