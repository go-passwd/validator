package validator

import (
	"fmt"
	"strings"
)

// ContainsOnly returns a ValidateFunc that check if password contains only selected chars
func ContainsOnly(chars string, customError error) ValidateFunc {
	return ValidateFunc(func(password string) error {
		for _, char := range []rune(password) {
			idx := strings.IndexFunc(chars, func(r rune) bool {
				return r == char
			})
			if idx == -1 {
				if customError != nil {
					return customError
				}
				return fmt.Errorf("the password must contains only %s", chars)
			}
		}
		return nil
	})
}
