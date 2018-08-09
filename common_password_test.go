package validator

import "fmt"

func ExampleCommonPassword() {
	passwordValidator := CommonPassword()
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("qaz123"))
	fmt.Println(passwordValidator("pa$$w0rd@"))
	// Output:
	// Password can't be a commonly used password
	// Password can't be a commonly used password
	// <nil>
}
