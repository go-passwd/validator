package validator_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-passwd/validator"
)

func TestStartsWith_emptyPassword(t *testing.T) {
	v := validator.StartsWith("abc", nil)
	assert.Error(t, v(""))
}

func ExampleStartsWith() {
	passwordValidator := validator.StartsWith("abc", nil)
	fmt.Println(passwordValidator("bui87j"))
	fmt.Println(passwordValidator("qwerty"))
	// Output:
	// <nil>
	// the password must start with one of: abc
}

func ExampleStartsWith_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.StartsWith("abc", err)
	fmt.Println(passwordValidator("bui87j"))
	fmt.Println(passwordValidator("qwerty"))
	// Output:
	// <nil>
	// custom error message
}
