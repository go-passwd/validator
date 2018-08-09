package validator

import (
	"errors"
	"fmt"
)

func ExampleRegex() {
	passwordValidator := Regex("^\\w+$", nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("pa$$w0rd"))
	// Output:
	// Password shouldn't match "^\w+$" pattern
	// <nil>
}

func ExampleRegex_customError() {
	err := errors.New("custom error message")
	passwordValidator := Regex("^\\w+$", err)
	fmt.Println(passwordValidator("password"))
	// Output:
	// custom error message
}
