package security

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash function that receives a string and returns a hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Verify function that compares a password and a hash and returns whether they are equal
func Verify(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}