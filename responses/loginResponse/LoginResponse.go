package loginresponse

import (
	userresponse "github.com/baskararestu/transfer-money/responses/userResponse"
)

type Login struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    LoginData `json:"data,omitempty"`
}

type LoginData struct {
	Token string                `json:"token"`
	User  userresponse.UserData `json:"user"`
}
