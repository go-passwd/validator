package validator_test

import (
	"errors"
	"fmt"

	"github.com/go-passwd/validator"
)

func ExampleNoop() {
	passwordValidator := validator.Noop(nil)
	fmt.Println(passwordValidator("password"))
	// Output:
	// <nil>
}

func ExampleNoop_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.Noop(err)
	fmt.Println(passwordValidator("password"))
	// Output:
	// custom error message
}
