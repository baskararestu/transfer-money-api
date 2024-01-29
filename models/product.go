package models

import (
	"github.com/google/uuid"
)

type Product struct {
	Id          string `gorm:"type:uuid;primaryKey" json:"id"`
	NameProduct string `gorm:"type:varchar(300)" json:"name_product"`
	Description string `gorm:"type:varchar(300)" json:"description"`
}

func NewProduct(name, description string) *Product {
	return &Product{
		Id:          uuid.New().String(),
		NameProduct: name,
		Description: description,
	}
}
