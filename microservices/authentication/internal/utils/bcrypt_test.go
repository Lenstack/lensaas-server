package utils

import "testing"

func TestHashPassword(t *testing.T) {
	bcrypt := NewBcrypt()
	hashedPassword, err := bcrypt.HashPassword("test")
	if err != nil {
		t.Error(err)
	}
	if hashedPassword == "" {
		t.Error("hashedPassword is empty")
	}
	t.Log(hashedPassword)
}

func TestComparePassword(t *testing.T) {
	bcrypt := NewBcrypt()
	hashedPassword, err := bcrypt.HashPassword("test")
	if err != nil {
		t.Error(err)
	}
	if hashedPassword == "" {
		t.Error("hashedPassword is empty")
	}

	err = bcrypt.ComparePassword(hashedPassword, "test")
	if err != nil {
		t.Error(err)
	}
	t.Log("password is correct")
}
