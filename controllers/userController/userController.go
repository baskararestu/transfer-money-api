package usercontroller

import (
	"net/http"

	db "enigma.com/learn-golang/database"
	"enigma.com/learn-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response := Response{
			Success: false,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	user.Id = uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Failed to hash password",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	user.Password = string(hashedPassword)

	if err := db.DB.Create(&user).Error; err != nil {
		response := Response{
			Success: false,
			Message: "Failed to create user",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	user.Password = "" // Remove password from response
	response := Response{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	}
	c.JSON(http.StatusOK, response)
}

func Update(c *gin.Context) {
	// Implement update logic
}

func Delete(c *gin.Context) {
	// Implement delete logic
}
