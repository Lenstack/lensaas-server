package utils

import "github.com/go-playground/validator/v10"

type ValidateError struct {
	FailedField string
	Tag         string
	Value       string
}

func Validate(i interface{}) (errors []*ValidateError) {
	validate := validator.New()
	err := validate.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			var element ValidateError
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
