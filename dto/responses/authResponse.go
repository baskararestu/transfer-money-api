package response

import (
	"time"

	"github.com/baskararestu/transfer-money/models"
)

// LoginResponse represents the response structure for login endpoint.
type LoginResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    LoginResponseData `json:"data,omitempty"`
}

// LoginResponseData represents the data structure within LoginResponse.
type LoginResponseData struct {
	Token string   `json:"token"`
	User  UserData `json:"user"`
}

// UserData represents the user data structure within LoginResponseData.
type UserData struct {
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewLoginResponse creates a new LoginResponse object.
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

// ErrorResponse represents the response structure for errors.
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// NewErrorResponse creates a new ErrorResponse object.
func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: message,
	}
}

// CreateUserResponse represents the response structure for user creation endpoint.
type CreateUserResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    *UserDetail `json:"data,omitempty"`
}

// UserDetail represents the user detail structure within CreateUserResponse.
type UserDetail struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewCreateUserResponse creates a new CreateUserResponse object.
func NewCreateUserResponse(user *models.User) *CreateUserResponse {
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
	}
}
