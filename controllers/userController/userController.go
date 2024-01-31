package usercontroller

import (
	"net/http"

	db "github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/models"
	mainresponse "github.com/baskararestu/transfer-money/responses/mainResponse"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var users []models.User

	db.DB.Find(&users)
	response := mainresponse.DataResponse{
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
			response := mainresponse.DataResponse{
				Success: false,
				Message: "User not found",
			}
			c.JSON(http.StatusNotFound, response)
			return
		default:
			response := mainresponse.DataResponse{
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
}

func Delete(c *gin.Context) {
}
