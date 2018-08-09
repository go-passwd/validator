package validator

import "fmt"

func ExampleNoop() {
	passwordValidator := Noop()
	fmt.Println(passwordValidator("password"))
	// Output:
	// <nil>
}
