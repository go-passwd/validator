package validator_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-passwd/validator"
)

func TestRatio(t *testing.T) {
	r := validator.Ratio("abc", "abc")
	assert.InDelta(t, 1.0, r, 1e-9) // Allow a small margin for floating-point precision

	r = validator.Ratio("abc", "def")
	assert.InDelta(t, 0.0, r, 1e-9) // Allow a small margin for floating-point precision

	r = validator.Ratio("abcd", "bcde")
	assert.InDelta(t, 0.75, r, 1e-9) // Allow a small margin for floating-point precision
}

func ExampleSimilarity() {
	passwordValidator := validator.Similarity([]string{"username", "username@example.com"}, nil, nil)
	fmt.Println(passwordValidator("username"))
	fmt.Println(passwordValidator("example"))
	similarity := 0.5
	passwordValidator = validator.Similarity([]string{"username", "username@example.com"}, &similarity, nil)
	fmt.Println(passwordValidator("username"))
	fmt.Println(passwordValidator("examplecom"))
	// Output:
	// The password is too similar to the username
	// <nil>
	// The password is too similar to the username
	// The password is too similar to the username@example.com
}

func ExampleSimilarity_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.Similarity([]string{"username", "username@example.com"}, nil, err)
	fmt.Println(passwordValidator("username"))
	// Output:
	// custom error message
}
