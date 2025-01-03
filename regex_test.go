package validator_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-passwd/validator"
)

func ExampleRegex() {
	passwordValidator := validator.Regex("^\\w+$", nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("pa$$w0rd"))
	// Output:
	// Password shouldn't match "^\w+$" pattern
	// <nil>
}

func ExampleRegex_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.Regex("^\\w+$", err)
	fmt.Println(passwordValidator("password"))
	// Output:
	// custom error message
}

func TestRegex_badRegexp(t *testing.T) {
	passwordValidator := validator.Regex("^\\q+$", nil)
	err := passwordValidator("password")
	assert.Error(t, err)
}
