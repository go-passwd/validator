package validator

import (
	"fmt"
)

func ExampleNoop() {
	passwordValidator := Noop()
	fmt.Println(passwordValidator("password"))
	// Output:
	// true
}

func ExampleMinLength() {
	passwordValidator := MinLength(5)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("pass"))
	fmt.Println(passwordValidator("passw"))
	// Output:
	// true
	// false
	// true
}

func ExampleMaxLength() {
	passwordValidator := MaxLength(5)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("pass"))
	fmt.Println(passwordValidator("passw"))
	// Output:
	// false
	// true
	// true
}

func ExampleContainsAtLeast() {
	passwordValidator := ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 4)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("PASSWORD"))
	fmt.Println(passwordValidator("passWORD"))
	// Output:
	// true
	// false
	// true
}
