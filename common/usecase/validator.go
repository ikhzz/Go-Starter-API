package usecase

import (
	"fmt"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func (g *commonUsecase) Validate(param interface{}) (error, []string) {
	err := g.v.Struct(param)
	errMsg := make([]string, 0)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			var reason string
			switch e.Tag() {
			case "email":
				reason = "must be a valid email"
			case "min":
				reason = fmt.Sprintf("must be at least %s characters", e.Param())
			case "alpha_num":
				reason = "must contain both letters and numbers"
			case "oneof":
				reason = fmt.Sprintf(`%s must be one of %s `, e.Field(), e.Param())
			default:
				reason = fmt.Sprintf(`%s is %s`, e.Field(), e.Tag())
			}
			errMsg = append(errMsg, reason)
		}
	}
	return err, errMsg
}

func containsAlphaNum(fl validator.FieldLevel) bool {
	hasLetter := false
	hasNumber := false

	// Check each character in the string
	for _, char := range fl.Field().String() {
		if unicode.IsLetter(char) {
			hasLetter = true
		}
		if unicode.IsDigit(char) {
			hasNumber = true
		}
		// If both letter and number are found, return true
		if hasLetter && hasNumber {
			return true
		}
	}
	return false
}
