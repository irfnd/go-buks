package utils

import (
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func ValidatorNew() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("postgresql_uri", postgresUri)
	validate.RegisterValidation("strong_password", strongPassword)
	return validate
}

func postgresUri(fl validator.FieldLevel) bool {
	uri := fl.Field().String()
	return strings.HasPrefix(uri, "postgres://") || strings.HasPrefix(uri, "postgresql://")
}

func strongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
			case unicode.IsUpper(char):
				hasUpper = true
			case unicode.IsLower(char):
				hasLower = true
			case unicode.IsDigit(char):
				hasNumber = true
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}