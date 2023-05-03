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
	Secret                          string
	ExpirationTimeAccess            time.Duration
	ExpirationTimeRefresh           time.Duration
	ExpirationTimeResetPassword     time.Duration
	ExpirationTimeEmailVerification time.Duration
}

func NewJwt(secret string, expirationAccess string, expirationRefresh string,
	expirationResetPassword string, expirationEmailVerification string) *Jwt {
	// Parse expiration time from string to time.Duration
	expirationTimeAccess, err := time.ParseDuration(expirationAccess)
	if err != nil {
		panic(err)
	}
	// Parse expiration time from string to time.Duration
	expirationTimeRefresh, err := time.ParseDuration(expirationRefresh)
	if err != nil {
		panic(err)
	}
	// Parse expiration time from string to time.Duration
	expirationTimeResetPassword, err := time.ParseDuration(expirationResetPassword)
	if err != nil {
		panic(err)
	}
	// Parse expiration time from string to time.Duration
	expirationTimeEmailVerification, err := time.ParseDuration(expirationEmailVerification)
	if err != nil {
		panic(err)
	}

	// Return Jwt instance with parsed expiration time initialized
	return &Jwt{
		Secret:                          secret,
		ExpirationTimeAccess:            expirationTimeAccess,
		ExpirationTimeRefresh:           expirationTimeRefresh,
		ExpirationTimeResetPassword:     expirationTimeResetPassword,
		ExpirationTimeEmailVerification: expirationTimeEmailVerification,
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
