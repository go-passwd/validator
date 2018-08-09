package validator

import (
	"fmt"
)

func ExampleRegex() {
	passwordValidator := Regex("^\\w+$")
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("pa$$w0rd"))
	// Output:
	// Password shouldn't match "^\w+$" pattern
	// <nil>
}
