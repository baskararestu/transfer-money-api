package authcontroller

import (
	"net/http"

	response "github.com/baskararestu/transfer-money/dto/responses"
	"github.com/baskararestu/transfer-money/models"
	"github.com/baskararestu/transfer-money/services"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

// CreateUser handles user creation requests.
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response := response.NewErrorResponse(err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := services.CheckExistingUser(user.Email); err == nil {
		response := response.NewErrorResponse("Email already exists")
		c.JSON(http.StatusConflict, response)
		return
	}

	user.Id = uuid.New().String()
	if err := services.CreateUserRecord(&user); err != nil {
		response := response.NewErrorResponse("Failed to create user")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	user.Password = "" // Remove password from response
	response := response.NewCreateUserResponse(&user)
	c.JSON(http.StatusOK, response)
}

// Login handles user login requests.
func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response := response.NewErrorResponse(err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := services.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Invalid email or password"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Invalid email or password"))
		return
	}

	token, err := services.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to generate token"))
		return
	}

	loginResponse := response.NewLoginResponse(token, user)
	c.JSON(http.StatusOK, loginResponse)
}
