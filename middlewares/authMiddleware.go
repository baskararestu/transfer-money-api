package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		if IsTokenBlacklisted(tokenString) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is blacklisted"})
			return
		}

		jwtSecret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, err := uuid.Parse(claims["user_id"].(string))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			c.Set("userID", userID)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}

func GetUserIDFromContext(c *gin.Context) (uuid.UUID, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return uuid.Nil, errors.New("user ID not found in context")
	}
	return userID.(uuid.UUID), nil
}

var (
	blacklistMap     = make(map[string]time.Time)
	blacklistMapLock sync.Mutex
)

const TokenBlacklistDuration = 24 * time.Hour

func BlacklistToken(tokenString string) error {
	blacklistMapLock.Lock()
	defer blacklistMapLock.Unlock()

	if _, exists := blacklistMap[tokenString]; exists {
		return errors.New("token already blacklisted")
	}

	expirationTime := time.Now().Add(TokenBlacklistDuration)

	blacklistMap[tokenString] = expirationTime

	return nil
}

func IsTokenBlacklisted(tokenString string) bool {
	blacklistMapLock.Lock()
	defer blacklistMapLock.Unlock()

	expirationTime, exists := blacklistMap[tokenString]
	if !exists {
		return false
	}

	if time.Now().After(expirationTime) {
		delete(blacklistMap, tokenString)
		return false
	}

	return true
}
