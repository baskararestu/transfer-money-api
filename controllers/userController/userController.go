package usercontroller

import (
	"net/http"

	db "github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Index(c *gin.Context) {
	var users []models.User

	db.DB.Find(&users)
	response := Response{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    users,
	}
	c.JSON(http.StatusOK, response)
}

func Show(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := Response{
				Success: false,
				Message: "User not found",
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
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User retrieved successfully", "user": user})
}

func Update(c *gin.Context) {
	// Implement update logic
}

func Delete(c *gin.Context) {
	// Implement delete logic
}
