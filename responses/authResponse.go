package response

import (
	"time"

	"github.com/baskararestu/transfer-money/models"
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

func NewLoginResponse(token string, user *models.User) LoginResponse {
	return LoginResponse{
		Success: true,
		Message: "Login successful",
		Data: LoginResponseData{
			Token: token,
			User: UserData{
				FullName:  user.FullName,
				Email:     user.Email,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		},
	}
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: message,
	}
}

type BankAccountDetail struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type CreateUserResponse struct {
	Success     bool               `json:"success"`
	Message     string             `json:"message"`
	Data        *UserDetail        `json:"data,omitempty"`
	BankAccount *BankAccountDetail `json:"bank_account,omitempty"`
}

type UserDetail struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCreateUserResponse(user *models.User, bankAccount *models.BankAccount) *CreateUserResponse {
	return &CreateUserResponse{
		Success: true,
		Message: "User created successfully",
		Data: &UserDetail{
			ID:        user.Id,
			FullName:  user.FullName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		BankAccount: &BankAccountDetail{
			ID:        bankAccount.ID.String(),
			UserID:    bankAccount.UserID.String(),
			Balance:   bankAccount.Balance,
			CreatedAt: bankAccount.CreatedAt,
			UpdatedAt: bankAccount.UpdatedAt,
		},
	}
}
