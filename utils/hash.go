package utils

import (
	"golang.org/x/crypto/bcrypt"
)
func HashPassword(password string) (string, error){
	// GenerateFromPassword will require byte slice
	// []byte(password) will convert string to byte
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}