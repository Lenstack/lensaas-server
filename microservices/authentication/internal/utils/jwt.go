package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type IJwt interface {
	GenerateToken(subject interface{}) (token string, err error)
	ValidateToken(token string) (subject interface{}, err error)
}

type Jwt struct {
	Secret                string
	ExpirationTimeAccess  time.Duration
	ExpirationTimeRefresh time.Duration
}

func NewJwt(secret string, expirationAccess string, expirationRefresh string) *Jwt {
	expirationTimeAccess, err := time.ParseDuration(expirationAccess)
	if err != nil {
		panic(err)
	}
	expirationTimeRefresh, err := time.ParseDuration(expirationRefresh)
	if err != nil {
		panic(err)
	}
	return &Jwt{
		Secret:                secret,
		ExpirationTimeAccess:  expirationTimeAccess,
		ExpirationTimeRefresh: expirationTimeRefresh,
	}
}

func (j *Jwt) GenerateToken(userId string, expiration time.Duration) (token string, err error) {
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(expiration).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Secret))
}

func (j *Jwt) ValidateToken(token string) (userId string, err error) {
	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	parsedToken, err := parser.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}
	return claims["id"].(string), nil
}