package utils

import (
	"testing"
)

type TestStruct struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func TestValidate(t *testing.T) {
	validationErrors := Validate(&TestStruct{
		Username: "test",
		Password: "test",
	})

	if len(validationErrors) > 0 {
		t.Error("validation errors")
	}

	t.Log(validationErrors)

}
