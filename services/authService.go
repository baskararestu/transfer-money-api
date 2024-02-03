package services

import (
	"errors"
	"os"
	"time"

	"github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckExistingUser(email string) error {
	var existingUser models.User
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUserRecordInTransaction(tx *gorm.DB, user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if err := tx.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func Login(email, password string) (string, *models.User, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	token, err := GenerateToken(user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func GenerateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	jwtSecret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(jwtSecret))
}
