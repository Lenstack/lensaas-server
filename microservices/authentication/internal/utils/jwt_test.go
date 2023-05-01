package utils

import (
	"testing"
)

// TestGenerateJWT ...
func TestGenerateJWT(t *testing.T) {
	jwt := NewJwt("secret", "1h", "24h")
	token, err := jwt.GenerateToken("test", jwt.ExpirationTimeAccess)
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Error("token is empty")
	}
	t.Log(token)
}

// TestValidateJWT ...
func TestValidateJWT(t *testing.T) {
	jwt := NewJwt("secret", "1h", "24h")
	token, err := jwt.GenerateToken("test", jwt.ExpirationTimeAccess)
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Error("token is empty")
	}
	t.Log(token)
	subject, err := jwt.ValidateToken(token)
	if err != nil {
		t.Error(err)
	}
	if subject != "test" {
		t.Error("subject is not equal to test")
	}
	t.Log(subject)
}
