package validator

import (
	"fmt"
	"strings"
)

// Similarity returns ValidateFunc that validate whether the password is sufficiently different from the attributes
//
// Attributes can be: user login, email, first name, last name, …
func Similarity(attributes []string, maxSimilarity *float64, customError error) ValidateFunc {
	if maxSimilarity == nil {
		maxSimilarity = new(float64)
		*maxSimilarity = 0.7
	}
	return ValidateFunc(func(password string) error {
		for idx := range attributes {
			if ratio(password, attributes[idx]) > *maxSimilarity {
				if customError != nil {
					return customError
				}
				return fmt.Errorf("The password is too similar to the %s", attributes[idx])
			}
		}
		return nil
	})
}

// Return a measure of the sequences' similarity (float in [0,1]).
//
// Where T is the total number of elements in both sequences, and M is the number of matches, this is 2.0*M / T.
// Note that this is 1 if the sequences are identical, and 0 if they have nothing in common.
func ratio(a, b string) float64 {
	t := float64(len(a) + len(b))
	m := 0.0
	sa := strings.Split(a, "")
	sb := strings.Split(b, "")
	for idx := range sb {
		aidx := strings.Index(a, sb[idx])
		if aidx > -1 {
			sa = append(sa[:aidx], sa[aidx+1:]...)
			a = strings.Join(sa, "")
			m = m + 1.0
		}
	}
	return 2.0 * m / t
}
