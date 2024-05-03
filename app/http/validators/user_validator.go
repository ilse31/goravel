package validators

import (
	"github.com/gookit/validate"
)

type UserValidator struct {
}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (r *UserValidator) Json(data map[string]interface{}) *validate.Validation {
	v := validate.New(data)
	r.validateFields(v, map[string]string{
		"name":                 "required",
		"avatar":               "required",
		"phoneNumber":          "required|phone",
		"email":                "required|email",
		"password":             "required",
		"passwordConfirmation": "required|same:password",
		"savings":              "required|numeric",
		"current":              "required|numeric",
		"credit":               "required|numeric",
		"role":                 "required",
		"status":               "required",
		"address":              "required",
		"city":                 "required",
		"state":                "required",
		"zip":                  "required",
		"country":              "required",
		"dateOfBirth":          "required|date",
	})
	return v
}

func (r *UserValidator) validateFields(v *validate.Validation, fields map[string]string) {
	for field, rule := range fields {
		v.StringRule(field, rule)
	}
}
