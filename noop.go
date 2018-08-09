package validator

// Noop returns a ValidateFunc that always truth
func Noop() ValidateFunc {
	return ValidateFunc(func(password string) error {
		return nil
	})
}
