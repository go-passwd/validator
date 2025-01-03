package validator_test

import (
	"errors"
	"fmt"

	"github.com/go-passwd/validator"
)

func ExampleStartsWith() {
	passwordValidator := validator.StartsWith("abc", nil)
	fmt.Println(passwordValidator("bui87j"))
	fmt.Println(passwordValidator("qwerty"))
	// Output:
	// <nil>
	// the password must starts with one of: abc
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
