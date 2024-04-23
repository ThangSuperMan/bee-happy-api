package auth

import (
	"testing"
)

func TestCreateJWT(t *testing.T) {
	secrete := []byte("secret")

	token, err := CreateJWT(secrete, 1)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Error("expected token to be not empty")
	}
}
