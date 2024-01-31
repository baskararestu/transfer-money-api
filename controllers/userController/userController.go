package usercontroller

import (
	"net/http"

	db "github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/models"
	mainresponse "github.com/baskararestu/transfer-money/responses/mainResponse"
	userService "github.com/baskararestu/transfer-money/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	users, err := userService.GetAllUsers()
	if err != nil {
		response := mainresponse.DefaultResponse{
			Success: false,
			Message: "Failed to retrieve users",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := mainresponse.DataResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    users,
	}
	c.JSON(http.StatusOK, response)
}

func Show(c *gin.Context) {
	id := c.Param("id")
	user, err := userService.GetUserByID(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := mainresponse.DefaultResponse{
				Success: false,
				Message: "User not found",
			}
			c.JSON(http.StatusNotFound, response)
			return
		default:
			response := mainresponse.DefaultResponse{
				Success: false,
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	response := mainresponse.DataResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    user,
	}
	c.JSON(http.StatusOK, response)
}

func Update(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := mainresponse.DefaultResponse{
				Success: false,
				Message: "User not found",
			}
			c.JSON(http.StatusNotFound, response)
			return
		default:
			response := mainresponse.DefaultResponse{
				Success: false,
				Message: err.Error(),
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userService.UpdateUser(&user); err != nil {
		response := mainresponse.DefaultResponse{
			Success: false,
			Message: "Failed to update user",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := mainresponse.DefaultResponse{
		Success: true,
		Message: "User updated successfully",
	}
	c.JSON(http.StatusOK, response)
}
