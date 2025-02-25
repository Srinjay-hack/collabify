package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key") // Store this in an environment variable for production

// Claims struct for JWT
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Set expiration time
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // Set the expiration time correctly
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateJWT validates the JWT token
func ValidateJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
