package validator

import (
	"fmt"
)

func ExampleNoop() {
	passwordValidator := Noop()
	fmt.Println(passwordValidator("password"))
	// Output:
	// <nil>
}

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

func ExampleContainsAtLeast() {
	passwordValidator := ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 4)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("PASSWORD"))
	fmt.Println(passwordValidator("passWORD"))
	// Output:
	// <nil>
	// Password must contains at least 4 chars from abcdefghijklmnopqrstuvwxyz
	// <nil>
}
