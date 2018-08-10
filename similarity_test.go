package validator

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRatio(t *testing.T) {
	r := ratio("abc", "abc")
	assert.Equal(t, 1.0, r)
	r = ratio("abc", "def")
	assert.Equal(t, 0.0, r)
	r = ratio("abcd", "bcde")
	assert.Equal(t, 0.75, r)
}

func ExampleSimilarity() {
	passwordValidator := Similarity([]string{"username", "username@example.com"}, nil, nil)
	fmt.Println(passwordValidator("username"))
	fmt.Println(passwordValidator("example"))
	similarity := 0.5
	passwordValidator = Similarity([]string{"username", "username@example.com"}, &similarity, nil)
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
	passwordValidator := Similarity([]string{"username", "username@example.com"}, nil, err)
	fmt.Println(passwordValidator("username"))
	// Output:
	// custom error message
}
