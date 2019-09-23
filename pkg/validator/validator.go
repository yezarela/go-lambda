package validator

import (
	"github.com/asaskevich/govalidator"
)

// RegisterCustomValidator ...
func RegisterCustomValidator(name string, validator func(string) bool) {
	govalidator.TagMap[name] = govalidator.Validator(validator)
}

// ValidateStruct ...
func ValidateStruct(d interface{}) error {

	_, err := govalidator.ValidateStruct(d)

	return err
}
