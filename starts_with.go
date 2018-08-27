package validator

import (
	"bytes"
	"fmt"
)

// StartsWith returns ValidateFunc that validate whether the password is starts with one of letter
func StartsWith(letters string, customError error) ValidateFunc {
	return ValidateFunc(func(password string) error {
		firstLetter := []byte(password)[0]
		idx := bytes.IndexByte([]byte(letters), firstLetter)
		if idx == -1 {
			if customError != nil {
				return customError
			}
			return fmt.Errorf("the password must starts with one of: %s", letters)
		}
		return nil
	})
}
