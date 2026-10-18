package api

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type StructValidator struct {
	validator *validator.Validate
}

func NewStructValidator() *StructValidator {
	v := validator.New(validator.WithRequiredStructEnabled())
	if err := v.RegisterValidation("nonblank", func(field validator.FieldLevel) bool {
		return strings.TrimSpace(field.Field().String()) != ""
	}); err != nil {
		panic(err)
	}
	return &StructValidator{validator: v}
}

func (v *StructValidator) Validate(out any) error {
	return v.validator.Struct(out)
}
