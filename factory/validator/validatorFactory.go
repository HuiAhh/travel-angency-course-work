package validator

import "github.com/go-playground/validator"

var validate *validator.Validate

func init() {
	//validate
	validate = validator.New()
	validate.RegisterValidation("is-password", func(fl validator.FieldLevel) bool {

		// service.UserService
		return fl.Field().Len() > 5
	})
}

//valildator
func GetValidator() *validator.Validate {
	return validate
}
