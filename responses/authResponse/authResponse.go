package authresponse

import (
	"github.com/baskararestu/transfer-money/models"
	bankresponse "github.com/baskararestu/transfer-money/responses/bankResponse"
	loginresponse "github.com/baskararestu/transfer-money/responses/loginResponse"
	userresponse "github.com/baskararestu/transfer-money/responses/userResponse"
)

func NewLoginResponse(token string, user *models.User) loginresponse.Login {
	return loginresponse.Login{
		Success: true,
		Message: "Login successful",
		Data: loginresponse.LoginData{
			Token: token,
			User: userresponse.UserData{
				FullName:  user.FullName,
				Email:     user.Email,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		},
	}
}

func NewCreateUserResponse(user *models.User, bankAccount *models.BankAccount) *userresponse.CreateUser {
	return &userresponse.CreateUser{
		Success: true,
		Message: "User created successfully",
		Data: &userresponse.UserDetail{
			ID:        user.Id.String(),
			FullName:  user.FullName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		BankAccount: &bankresponse.BankAccountDetail{
			ID:            bankAccount.ID.String(),
			UserID:        bankAccount.UserID.String(),
			AccountNumber: bankAccount.AccountNumber,
			Balance:       bankAccount.Balance,
			CreatedAt:     bankAccount.CreatedAt,
			UpdatedAt:     bankAccount.UpdatedAt,
		},
	}
}
