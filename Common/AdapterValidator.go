package Common

import "gopkg.in/go-playground/validator.v9"

type AdapterValidator struct {
	validator *validator.Validate
}

func (cv *AdapterValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewCustomValidator() *AdapterValidator {
	cusValidator := new(AdapterValidator)
	cusValidator.validator = validator.New()
	return cusValidator
}
