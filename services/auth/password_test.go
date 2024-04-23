package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("expected hash to be not empty")
	}

	if hash == "password" {
		t.Error("expected hash to be different from password")
	}
}

func TestComparePasswords(t *testing.T) {
	expectedPassword := "password"
	hash, err := HashPassword(expectedPassword)
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePassword(hash, []byte(expectedPassword)) {
		t.Errorf("expected password to match hash")
	}

	if ComparePassword(hash, []byte("notpassword")) {
		t.Errorf("expected password to not match hash")
	}
}