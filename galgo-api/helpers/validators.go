package helpers

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func Validate(value interface{}) error {
	validate := validator.New()
	err := validate.Struct(value)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
