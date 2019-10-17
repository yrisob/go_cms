package custom_validator

import "gopkg.in/go-playground/validator.v9"

type CustomValidator struct {
	validator *validator.Validate
}

//NewValidator create special  new validation for project needed
func NewValidator() *CustomValidator {
	v := validator.New()
	_ = v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "" || len(fl.Field().String()) == 11
	})

	_ = v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "" || len(fl.Field().String()) >= 6
	})

	return &CustomValidator{validator: v}
}

//Validate 	function for validate dto objects
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
