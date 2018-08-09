package validator

import (
	"errors"
	"fmt"
)

func ExampleContainsAtLeast() {
	passwordValidator := ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 4, nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("PASSWORD"))
	fmt.Println(passwordValidator("passWORD"))
	// Output:
	// <nil>
	// Password must contains at least 4 chars from abcdefghijklmnopqrstuvwxyz
	// <nil>
}

func ExampleContainsAtLeast_customError() {
	err := errors.New("custom error message")
	passwordValidator := ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 4, err)
	fmt.Println(passwordValidator("PASSWORD"))
	// Output:
	// custom error message
}
