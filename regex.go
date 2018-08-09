package validator

import (
	"fmt"
	"regexp"
)

// Regex returns ValidateFunc that check if password match regexp pattern
func Regex(pattern string) ValidateFunc {
	return ValidateFunc(func(password string) error {
		matched, err := regexp.MatchString(pattern, password)
		if err != nil {
			return err
		}
		if matched {
			return fmt.Errorf("Password shouldn't match \"%s\" pattern", pattern)
		}
		return nil
	})
}
