package validator_test

import (
	"errors"
	"fmt"

	"github.com/go-passwd/validator"
)

func ExampleContainsOnly() {
	passwordValidator := validator.ContainsOnly("abcdefghijklmnopqrstuvwxyz", nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("password0"))
	fmt.Println(passwordValidator("passWORD"))
	// Output:
	// <nil>
	// The password must contain only abcdefghijklmnopqrstuvwxyz
	// The password must contain only abcdefghijklmnopqrstuvwxyz
}

func ExampleContainsOnly_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.ContainsOnly("abcdefghijklmnopqrstuvwxyz", err)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("password0"))
	fmt.Println(passwordValidator("passWORD"))
	// Output:
	// <nil>
	// custom error message
	// custom error message
}
