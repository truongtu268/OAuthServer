package NewMiddleware

import (
	"errors"
	"strings"
	"reflect"
)

type ValidatorLocate struct {
	listValidator map[string]IValidator
}

func (validators *ValidatorLocate) AddValidator(validator IValidator) error {
	var NameValidator = strings.Replace(reflect.TypeOf(validator).String(),"*NewMiddleware.","",-1)
	_, ok:=validators.listValidator[NameValidator]
	if ok {
		return errors.New("This key exist in Validator Locate")
	}
	validators.listValidator[NameValidator] = validator
	return nil
}

func (validators *ValidatorLocate) GetValidator(name string) (error, IValidator) {
	validator,ok := validators.listValidator[name]
	if ok {
		return nil, validator
	}
	return errors.New("This service doesn't exist in Service locator"), nil
}

func NewValidatorLocation() *ValidatorLocate {
	validators := new(ValidatorLocate)
	validators.listValidator = make(map[string]IValidator)
	validators.AddValidator(new(ValidateHeaderToken))
	validators.AddValidator(new(ValidateTest))
	return validators
}
