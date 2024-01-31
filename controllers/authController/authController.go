package authcontroller

import (
	"net/http"

	"github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/models"
	response "github.com/baskararestu/transfer-money/responses"
	"github.com/baskararestu/transfer-money/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var user models.User
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
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

	userID := uuid.New().String()

	user.Id = userID

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		response := response.NewErrorResponse("Failed to hash password")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	user.Password = string(hashedPassword)

	if err := services.CreateUserRecordInTransaction(tx, &user); err != nil {
		tx.Rollback()
		response := response.NewErrorResponse("Failed to create user")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	accountNumber := services.GenerateAccountNumber()

	bankAccountID, err := services.CreateBankAccountForUser(tx, userID, accountNumber)
	if err != nil {
		tx.Rollback()
		response := response.NewErrorResponse("Failed to create bank account")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	tx.Commit()

	bankAccount, err := services.GetBankAccountByID(bankAccountID)
	if err != nil {
		response := response.NewErrorResponse("Failed to retrieve bank account details")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	user.Password = ""
	response := response.NewCreateUserResponse(&user, bankAccount)
	c.JSON(http.StatusOK, response)
}

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

	token, user, err := services.Login(req.Email, req.Password)
	if err != nil {
		response := response.NewErrorResponse(err.Error())
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	loginResponse := response.NewLoginResponse(token, user)
	c.JSON(http.StatusOK, loginResponse)
}
