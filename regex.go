package validator

import (
	"fmt"
	"regexp"
)

// Regex returns ValidateFunc that check if password match regexp pattern.
func Regex(pattern string, customError error) ValidateFunc {
	compiled, compileErr := regexp.Compile(pattern)
	return ValidateFunc(func(password string) error {
		if compileErr != nil {
			return compileErr
		}
		if compiled.MatchString(password) {
			if customError != nil {
				return customError
			}
			return fmt.Errorf("Password shouldn't match \"%s\" pattern", pattern)
		}
		return nil
	})
}
