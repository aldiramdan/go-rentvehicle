package libs

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil

}

func CheckPassword(bodyPass, dbPass string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(bodyPass), []byte(dbPass))

	if err != nil {
		return false
	}

	return true
}
