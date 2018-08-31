package validator

import (
	"errors"
	"fmt"
)

func ExampleUnique() {
	passwordValidator := Unique(nil)
	fmt.Println(passwordValidator("bui87j"))
	fmt.Println(passwordValidator("qwerte"))
	// Output:
	// <nil>
	// the password must contains unique chars
}

func ExampleUnique_customError() {
	err := errors.New("custom error message")
	passwordValidator := Unique(err)
	fmt.Println(passwordValidator("bui87j"))
	fmt.Println(passwordValidator("qwerte"))
	// Output:
	// <nil>
	// custom error message
}
