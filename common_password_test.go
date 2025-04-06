package validator_test

import (
	"errors"
	"fmt"

	"github.com/go-passwd/validator"
)

func ExampleCommonPassword() {
	passwordValidator := validator.CommonPassword(nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("qaz123"))
	fmt.Println(passwordValidator("pa$$w0rd@"))
	fmt.Println(passwordValidator("Mod7tygrysow"))
	// Output:
	// Password can't be a commonly used password
	// Password can't be a commonly used password
	// <nil>
	// Password can't be a commonly used password
}

func ExampleCommonPassword_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.CommonPassword(err)
	fmt.Println(passwordValidator("password"))
	// Output:
	// custom error message
}
