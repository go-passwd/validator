package validator_test

import (
	"errors"
	"fmt"

	"github.com/go-passwd/validator"
)

func ExampleUnique() {
	passwordValidator := validator.Unique(nil)
	fmt.Println(passwordValidator("bui87j"))
	fmt.Println(passwordValidator("qwerte"))
	// Output:
	// <nil>
	// the password must contains unique chars
}

func ExampleUnique_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.Unique(err)
	fmt.Println(passwordValidator("bui87j"))
	fmt.Println(passwordValidator("qwerte"))
	// Output:
	// <nil>
	// custom error message
}
