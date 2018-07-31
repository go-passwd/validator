// Package validator is a set of functions to validate passwords
package validator

import "strings"

// ValidateFunc defines a function to validate
type ValidateFunc func(password string) bool

// Noop returns a ValidateFunc that always truth
func Noop() ValidateFunc {
	return ValidateFunc(func(password string) bool {
		return true
	})
}

// MinLength returns a ValidateFunc that check if password length is not lower that "length"
func MinLength(length int) ValidateFunc {
	return ValidateFunc(func(password string) bool {
		return len(password) >= length
	})
}

// MaxLength returns a ValidateFunc that check if password length is not greater that "length"
func MaxLength(length int) ValidateFunc {
	return ValidateFunc(func(password string) bool {
		return len(password) <= length
	})
}

// ContainsAtLeast returns a ValidateFunc that count occurrences of a chars and compares it with required value
func ContainsAtLeast(chars string, occurrences int) ValidateFunc {
	return ValidateFunc(func(password string) bool {
		cnt := 0
		aPassword := strings.Split(password, "")
		for _, char := range strings.Split(chars, "") {
			for _, pChar := range aPassword {
				if char == pChar {
					cnt++
				}
			}
		}
		return cnt >= occurrences
	})
}
