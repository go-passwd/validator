package validator

import (
	"fmt"
	"strings"
)

// ContainsAtLeast returns a ValidateFunc that count occurrences of a chars and compares it with required value.
func ContainsAtLeast(chars string, occurrences int, customError error) ValidateFunc {
	return ValidateFunc(func(password string) error {
		cnt := 0
		for _, char := range chars {
			cnt += strings.Count(password, string(char))
		}
		if cnt < occurrences {
			if customError != nil {
				return customError
			}
			return fmt.Errorf("Password must contain at least %d chars from %s", occurrences, chars)
		}
		return nil
	})
}
