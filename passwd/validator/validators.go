// Package validator is a set of functions to validate passwords
package validator

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
