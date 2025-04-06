package validator

// Noop returns a ValidateFunc that always return custom error.
func Noop(customError error) ValidateFunc {
	return ValidateFunc(func(password string) error {
		return customError
	})
}
