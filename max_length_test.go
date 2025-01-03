package validator_test

import (
	"errors"
	"fmt"

	"github.com/go-passwd/validator"
)

func ExampleMaxLength() {
	passwordValidator := validator.MaxLength(5, nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("pass"))
	fmt.Println(passwordValidator("passw"))
	// Output:
	// Password length must be not greater that 5 chars
	// <nil>
	// <nil>
}

func ExampleMaxLength_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.MaxLength(5, err)
	fmt.Println(passwordValidator("password"))
	// Output:
	// custom error message
}
