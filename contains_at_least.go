package validator

import (
	"fmt"
	"strings"
)

// ContainsAtLeast returns a ValidateFunc that count occurrences of a chars and compares it with required value
func ContainsAtLeast(chars string, occurrences int) ValidateFunc {
	return ValidateFunc(func(password string) error {
		cnt := 0
		aPassword := strings.Split(password, "")
		for _, char := range strings.Split(chars, "") {
			for _, pChar := range aPassword {
				if char == pChar {
					cnt++
				}
			}
		}
		if cnt < occurrences {
			return fmt.Errorf("Password must contains at least %d chars from %s", occurrences, chars)
		}
		return nil
	})
}
