package services

import (
	"github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/models"
)

func UpdateUser(user *models.User) error {
	if err := database.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id string) error {
	if err := database.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
