package validator

import (
	"errors"
	"fmt"
)

func ExampleNoop() {
	passwordValidator := Noop(nil)
	fmt.Println(passwordValidator("password"))
	// Output:
	// <nil>
}

func ExampleNoop_customError() {
	err := errors.New("custom error message")
	passwordValidator := Noop(err)
	fmt.Println(passwordValidator("password"))
	// Output:
	// custom error message
}
