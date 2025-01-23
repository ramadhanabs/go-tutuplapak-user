package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain text password.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

// CheckPasswordHash compares a hashed password with its plain text counterpart.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
