package productcontroller

import (
	"net/http"

	db "enigma.com/learn-golang/database"
	"enigma.com/learn-golang/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Response represents the structure of the response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Index(c *gin.Context) {
	var products []models.Product

	db.DB.Find(&products)
	response := Response{
		Success: true,
		Message: "Products retrieved successfully",
		Data:    products,
	}
	c.JSON(http.StatusOK, response)
}

func Show(c *gin.Context) {
	var products models.Product
	id := c.Param("id")

	if err := db.DB.First(&products, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := Response{
				Success: false,
				Message: "Product not found",
			}
			c.JSON(http.StatusNotFound, response)
			return
		default:
			response := Response{
				Success: false,
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Products retrieved successfully", "products": products})
}

func Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		response := Response{
			Success: false,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
	}

	db.DB.Create(&product)
	response := Response{
		Success: true,
		Message: "Product created successfully",
		Data:    product,
	}
	c.JSON(http.StatusOK, response)
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
