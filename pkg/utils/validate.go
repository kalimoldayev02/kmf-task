package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	validate := validator.New()

	_ = validate.RegisterValidation("currency_date", func(fl validator.FieldLevel) bool {
		date := fl.Field().String()
		// валидация для currency date
		match, _ := regexp.MatchString(`^\d{2}-\d{2}-\d{4}$`, date)
		return match
	})

	return validate
}
