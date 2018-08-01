package passwd

import (
	"errors"

	"github.com/tomi77/go-passwd/passwd/validator"
)

// Validator represents set of password validators
type Validator []validator.ValidateFunc

// NewValidator return new instance of Validator
func NewValidator(vfunc ...validator.ValidateFunc) *Validator {
	v := Validator{}
	v = append(v, vfunc...)
	return &v
}

// Validate the password
func (v *Validator) Validate(password string) error {
	if len(*v) == 0 {
		return errors.New("Validator must contains at least 1 validator function")
	}
	for _, passwordValidator := range *v {
		err := passwordValidator(password)
		if err != nil {
			return err
		}
	}
	return nil
}
