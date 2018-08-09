package validator

import (
	"errors"
	"fmt"
)

func ExampleMinLength() {
	passwordValidator := MinLength(5, nil)
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
	passwordValidator := MinLength(5, err)
	fmt.Println(passwordValidator("pass"))
	// Output:
	// custom error message
}
