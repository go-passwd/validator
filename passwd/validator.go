package passwd

import "github.com/tomi77/go-passwd/passwd/validator"

// Validator represents set of password validators
type Validator []validator.ValidateFunc

// Validate the password
func (v *Validator) Validate(password string) bool {
	if len(*v) == 0 {
		return false
	}
	for _, passwordValidator := range *v {
		if !passwordValidator(password) {
			return false
		}
	}
	return true
}
