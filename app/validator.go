package app

import "github.com/go-playground/validator/v10"

// NewValidator creates and returns a new instance of the validator.
func NewValidator() *validator.Validate {
	return validator.New()
}
