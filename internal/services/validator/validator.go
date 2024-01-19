package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/ttacon/libphonenumber"
)

func CreateValidator() *validator.Validate {
	var validate = validator.New()
	validate.RegisterValidation("phone", validateCustomPhone)
	return validate
}

func validateCustomPhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	_, err := libphonenumber.Parse(phone, "RU")
	return err == nil
}
