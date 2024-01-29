package authcontroller

import (
	"net/http"
	"time"

	db "enigma.com/learn-golang/database"
	"enigma.com/learn-golang/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    LoginResponseData `json:"data,omitempty"`
}

type LoginResponseData struct {
	Token string   `json:"token"`
	User  UserData `json:"user"`
}

type UserData struct {
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, LoginResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	var user models.User
	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, LoginResponse{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, LoginResponse{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, LoginResponse{
			Success: false,
			Message: "Failed to generate token",
		})
		return
	}

	userData := UserData{
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	response := LoginResponse{
		Success: true,
		Message: "Login successful",
		Data: LoginResponseData{
			Token: tokenString,
			User:  userData,
		},
	}

	c.JSON(http.StatusOK, response)
}
