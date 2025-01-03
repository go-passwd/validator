package validator_test

import (
	"errors"
	"fmt"

	"github.com/go-passwd/validator"
)

func ExampleMinLength() {
	passwordValidator := validator.MinLength(5, nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("pass"))
	fmt.Println(passwordValidator("passw"))
	// Output:
	// <nil>
	// Password length must be not lower that 5 chars
	// <nil>
}

func ExampleMinLength_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.MinLength(5, err)
	fmt.Println(passwordValidator("pass"))
	// Output:
	// custom error message
}
