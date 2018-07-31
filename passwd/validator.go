package passwd

import "github.com/tomi77/go-passwd/passwd/validator"

// Validator represents set of password validators
type Validator []validator.ValidateFunc

// Validate the password
func (v *Validator) Validate(password string) error {
	if len(*v) == 0 {
		return nil
	}
	for _, passwordValidator := range *v {
		err := passwordValidator(password)
		if err != nil {
			return err
		}
	}
	return nil
}
