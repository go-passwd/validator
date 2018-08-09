package validator

import "fmt"

func ExampleMaxLength() {
	passwordValidator := MaxLength(5)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("pass"))
	fmt.Println(passwordValidator("passw"))
	// Output:
	// Password length must be not greater that 5 chars
	// <nil>
	// <nil>
}
