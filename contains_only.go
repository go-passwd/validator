package validator

import (
	"bytes"
	"fmt"
	"strings"
)

// ContainsOnly returns a ValidateFunc that check if password contains only selected chars
func ContainsOnly(chars string, customError error) ValidateFunc {
	return ValidateFunc(func(password string) error {
		bChars := []byte(chars)
		for _, char := range strings.Split(password, "") {
			bChar := []byte(char)[0]
			idx := bytes.IndexByte(bChars, bChar)
			if idx == -1 {
				if customError != nil {
					return customError
				}
				return fmt.Errorf("the password must contains only %s", chars)
			}
		}
		return nil
	})
}
