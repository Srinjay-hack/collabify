package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(bytes), nil
}

// ComparePassword compares a hashed password with a plaintext password
func ComparePassword(hashedPassword, plainPassword string) error {
	// Compare the hashed password with the provided plaintext password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return err // Returns an error if passwords don't match
	}
	return nil
}
