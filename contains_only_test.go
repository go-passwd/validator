package validator

import (
	"errors"
	"fmt"
)

func ExampleContainsOnly() {
	passwordValidator := ContainsOnly("abcdefghijklmnopqrstuvwxyz", nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("password0"))
	fmt.Println(passwordValidator("passWORD"))
	// Output:
	// <nil>
	// the password must contains only abcdefghijklmnopqrstuvwxyz
	// the password must contains only abcdefghijklmnopqrstuvwxyz
}

func ExampleContainsOnly_customError() {
	err := errors.New("custom error message")
	passwordValidator := ContainsOnly("abcdefghijklmnopqrstuvwxyz", err)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("password0"))
	fmt.Println(passwordValidator("passWORD"))
	// Output:
	// <nil>
	// custom error message
	// custom error message
}
