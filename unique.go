package validator

import (
	"errors"
)

// Unique returns ValidateFunc that validate whether the password has only unique chars.
func Unique(customError error) ValidateFunc {
	return ValidateFunc(func(password string) error {
		seen := make(map[rune]struct{})
		for _, r := range password {
			if _, exists := seen[r]; exists {
				if customError != nil {
					return customError
				}
				return errors.New("the password must contain unique chars")
			}
			seen[r] = struct{}{}
		}
		return nil
	})
}
