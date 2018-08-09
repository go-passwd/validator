package validator

import "fmt"

func ExampleMinLength() {
	passwordValidator := MinLength(5)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("pass"))
	fmt.Println(passwordValidator("passw"))
	// Output:
	// <nil>
	// Password length must be not lower that 5 chars
	// <nil>
}
