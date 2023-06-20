package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validations = map[string]func(validator.FieldLevel) bool{
	"is-title-ok": ValidateVideoTitle,
}

// example validation, title should not contain the word ass
func ValidateVideoTitle(field validator.FieldLevel) bool {
	s := field.Field().String()
	arr := strings.Split(s, " ")
	for _, word := range arr {
		if word == "ass" {
			return false
		}
	}
	return true
}
