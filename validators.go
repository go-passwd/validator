package validator

import (
	"fmt"
	"strings"
)

// ValidateFunc defines a function to validate
type ValidateFunc func(password string) error

// Noop returns a ValidateFunc that always truth
func Noop() ValidateFunc {
	return ValidateFunc(func(password string) error {
		return nil
	})
}

// MinLength returns a ValidateFunc that check if password length is not lower that "length"
func MinLength(length int) ValidateFunc {
	return ValidateFunc(func(password string) error {
		if len(password) < length {
			return fmt.Errorf("Password length must be not lower that %d chars", length)
		}
		return nil
	})
}

// MaxLength returns a ValidateFunc that check if password length is not greater that "length"
func MaxLength(length int) ValidateFunc {
	return ValidateFunc(func(password string) error {
		if len(password) > length {
			return fmt.Errorf("Password length must be not greater that %d chars", length)
		}
		return nil
	})
}

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
