package validator

import "fmt"

// MaxLength returns a ValidateFunc that check if password length is not greater that "length"
func MaxLength(length int) ValidateFunc {
	return ValidateFunc(func(password string) error {
		if len(password) > length {
			return fmt.Errorf("Password length must be not greater that %d chars", length)
		}
		return nil
	})
}
