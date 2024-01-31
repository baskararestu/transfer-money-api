package userresponse

import (
	"time"

	bankresponse "github.com/baskararestu/transfer-money/responses/bankResponse"
)

type UserData struct {
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UserDetail struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type CreateUser struct {
	Success     bool                            `json:"success"`
	Message     string                          `json:"message"`
	Data        *UserDetail                     `json:"data,omitempty"`
	BankAccount *bankresponse.BankAccountDetail `json:"bank_account,omitempty"`
}
