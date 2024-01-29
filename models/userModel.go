package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        string    `gorm:"primaryKey" json:"id"`
	FullName  string    `json:"full_name" validate:"required,min=2,max=100"`
	Password  string    `json:"password" validate:"required,min=6"`
	Email     string    `json:"email" validate:"email,required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(fullName, lastName, password, email, phone, userType string) *User {
	return &User{
		Id:        uuid.New().String(),
		FullName:  fullName,
		Password:  password,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
