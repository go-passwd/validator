package validator

import "fmt"

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
