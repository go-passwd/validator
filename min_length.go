package validator

import (
	"fmt"
	"unicode/utf8"
)

// MinLength returns a ValidateFunc that check if password length is not lower that "length".
func MinLength(length int, customError error) ValidateFunc {
	return ValidateFunc(func(password string) error {
		if utf8.RuneCountInString(password) < length {
			if customError != nil {
				return customError
			}
			return fmt.Errorf("Password length must be not lower that %d chars", length)
		}
		return nil
	})
}
