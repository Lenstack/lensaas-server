package utils

import "golang.org/x/crypto/bcrypt"

type IBcryptManager interface {
	HashPassword(password string) (string, error)
	CompareHashedPassword(hashedPassword string, password string) (bool, error)
}

type BcryptManager struct {
}

func (bm *BcryptManager) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func (bm *BcryptManager) CompareHashedPassword(hashedPassword string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
