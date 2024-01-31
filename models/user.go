package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FullName  string    `json:"full_name" validate:"required,min=2,max=100"`
	Password  string    `json:"password" validate:"required,min=6"`
	Email     string    `json:"email" validate:"email,required" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
