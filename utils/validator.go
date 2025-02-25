package utils

import (
	"regexp"
	"unicode"
)

// IsValidEmail checks if the email is valid
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// IsValidPassword checks the strength of the password
func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, ch := range password {
		if unicode.IsUpper(ch) {
			hasUpper = true
		} else if unicode.IsLower(ch) {
			hasLower = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}
