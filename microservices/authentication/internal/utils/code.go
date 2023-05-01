package utils

import (
	cryptorand "crypto/rand"
	"encoding/base64"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type ICode interface {
	GenerateDigitCode(digit int) string
	GenerateStateOauth() (state string, expiration time.Time)
	GenerateCsrfToken() string
}

type Code struct {
}

func NewCode() *Code {
	return &Code{}
}

func (c *Code) GenerateDigitCode(digit int) string {
	min := int(math.Pow10(digit - 1))
	max := int(math.Pow10(digit)) - 1
	code := rand.Intn(max-min+1) + min
	return strconv.Itoa(code)
}

func (c *Code) GenerateStateOauth() (state string, expiration int64) {
	expiration = time.Now().Add(2 * time.Minute).Unix()
	token := make([]byte, 32)
	if _, err := cryptorand.Read(token); err != nil {
		panic(err)
	}
	state = base64.URLEncoding.EncodeToString(token)
	return state, expiration
}

func (c *Code) GenerateCsrfToken() string {
	token := make([]byte, 32)
	if _, err := cryptorand.Read(token); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(token)
}
