package utils

import "golang.org/x/crypto/bcrypt"

type IBcrypt interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}

type Bcrypt struct {
}

func NewBcrypt() *Bcrypt {
	return &Bcrypt{}
}

func (b *Bcrypt) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func (b *Bcrypt) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
