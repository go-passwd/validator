package validator

import (
	"errors"
	"fmt"
)

func ExampleCommonPassword() {
	passwordValidator := CommonPassword(nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("qaz123"))
	fmt.Println(passwordValidator("pa$$w0rd@"))
	// Output:
	// Password can't be a commonly used password
	// Password can't be a commonly used password
	// <nil>
}

func ExampleCommonPassword_customError() {
	err := errors.New("custom error message")
	passwordValidator := CommonPassword(err)
	fmt.Println(passwordValidator("password"))
	// Output:
	// custom error message
}
