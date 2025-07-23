package util

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data interface{}) error {
	validate := validator.New()

	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	switch validationError.Tag() {
	case "required":
		return errors.New(validationError.StructField() + " is required")
	case "max":
		return errors.New(validationError.StructField() + " must be less than " + validationError.Param())
	case "min":
		return errors.New(validationError.StructField() + " must be greater than " + validationError.Param())
	case "email":
		return errors.New(validationError.StructField() + " must be a valid email address")
	}

	return nil
}
