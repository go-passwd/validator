package validator

import "fmt"

// MinLength returns a ValidateFunc that check if password length is not lower that "length"
func MinLength(length int) ValidateFunc {
	return ValidateFunc(func(password string) error {
		if len(password) < length {
			return fmt.Errorf("Password length must be not lower that %d chars", length)
		}
		return nil
	})
}
