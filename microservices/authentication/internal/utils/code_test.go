package utils

import (
	"testing"
)

func TestGenerateDigitCode(t *testing.T) {
	code := NewCode()
	digitCode := code.GenerateDigitCode(5)
	if digitCode == "" {
		t.Error("digitCode is empty")
	}
	t.Log(digitCode)
}

func TestGenerateStateOauth(t *testing.T) {
	code := NewCode()
	state, expiration := code.GenerateStateOauth()
	if state == "" {
		t.Error("state is empty")
	}
	if expiration == 0 {
		t.Error("expiration is empty")
	}
	t.Log(state, expiration)
}

func TestGenerateCsrfToken(t *testing.T) {
	code := NewCode()
	csrfToken := code.GenerateCsrfToken()
	if csrfToken == "" {
		t.Error("csrfToken is empty")
	}
	t.Log(csrfToken)
}
