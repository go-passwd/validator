package validator

import (
	"fmt"
	"regexp"
)

// Regex returns ValidateFunc that check if password match regexp pattern.
func Regex(pattern string, customError error) ValidateFunc {
	return ValidateFunc(func(password string) error {
		matched, err := regexp.MatchString(pattern, password)
		if err != nil {
			return err
		}
		if matched {
			if customError != nil {
				return customError
			}
			return fmt.Errorf("password shouldn't match \"%s\" pattern", pattern)
		}
		return nil
	})
}
