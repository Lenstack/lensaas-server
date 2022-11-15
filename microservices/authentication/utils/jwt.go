package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type IJwtManager interface {
	GenerateToken(userId string) (token string, err error)
	ValidateToken(token string) (userId string, err error)
}

type JwtManager struct {
	secret     string
	expiration string
}

func NewJwtManager(secret string, expiration string) *JwtManager {
	return &JwtManager{secret: secret, expiration: expiration}
}

func (jm *JwtManager) GenerateToken(userId string) (token string, err error) {
	expirationTime, err := time.ParseDuration(jm.expiration)
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(expirationTime).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jm.secret))
}

func (jm *JwtManager) ValidateToken(token string) (userId string, err error) {
	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	parsedToken, err := parser.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jm.secret), nil
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
