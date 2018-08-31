package validator

import (
	"errors"
	"strings"
)

// Unique returns ValidateFunc that validate whether the password has only unique chars
func Unique(customError error) ValidateFunc {
	return ValidateFunc(func(password string) error {
		runes := []rune(password)
		for idx := range runes {
			if strings.LastIndexFunc(password, func(r rune) bool {
				return r == runes[idx]
			}) != idx {
				if customError != nil {
					return customError
				}
				return errors.New("the password must contains unique chars")
			}
		}
		return nil
	})
}
