package services

import (
	"time"

	"github.com/baskararestu/transfer-money/database"
	"github.com/baskararestu/transfer-money/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// CheckExistingUser checks if a user with the given email already exists.
func CheckExistingUser(email string) error {
	var existingUser models.User
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err != nil {
		return err // User does not exist, return nil error
	}
	return nil // User exists, return non-nil error
}

// CreateUserRecord creates a new user record in the database.
func CreateUserRecord(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	if err := database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByEmail retrieves a user record from the database by email.
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GenerateToken generates a JWT token for authentication.
func GenerateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString([]byte("secret"))
}
