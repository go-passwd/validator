package validator

import (
	"fmt"
	"strings"
)

// Similarity returns ValidateFunc that validate whether the password is sufficiently different from the attributes
//
// Attributes can be: user login, email, first name, last name, ….
func Similarity(attributes []string, maxSimilarity *float64, customError error) ValidateFunc {
	if maxSimilarity == nil {
		maxSimilarity = new(float64)
		*maxSimilarity = 0.7
	}
	return ValidateFunc(func(password string) error {
		for idx := range attributes {
			if Ratio(password, attributes[idx]) > *maxSimilarity {
				if customError != nil {
					return customError
				}
				return fmt.Errorf("The password is too similar to the %s", attributes[idx])
			}
		}
		return nil
	})
}

// Ratio returns a measure of the sequences' similarity (float in [0,1]).
//
// Where T is the total number of elements in both sequences, and M is the number of matches, this is 2.0*M / T.
// Note that this is 1 if the sequences are identical, and 0 if they have nothing in common.
func Ratio(pass, attr string) float64 {
	totalChars := float64(len(pass) + len(attr))
	matches := 0.0

	for _, char := range attr {
		if strings.Contains(pass, string(char)) {
			pass = strings.Replace(pass, string(char), "", 1)
			matches++
		}
	}

	return 2.0 * matches / totalChars
}
