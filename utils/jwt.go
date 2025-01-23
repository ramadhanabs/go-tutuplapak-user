package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateJWT generates a new JWT token.
func GenerateJWT(identifier, secret string, expiryHours int) (string, error) {
	claims := jwt.MapClaims{
		"identifier": identifier,
		"exp":        time.Now().Add(time.Hour * time.Duration(expiryHours)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
