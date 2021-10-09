package validate

import validator2 "github.com/go-playground/validator/v10"

var (
	validator = validator2.New()
)

func Validate(data interface{}) error {
	return validator.Struct(data)
}
