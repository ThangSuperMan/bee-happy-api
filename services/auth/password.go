package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hased string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hased), plain)
	return err == nil
}
