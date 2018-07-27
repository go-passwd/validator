package util

import (
	"fmt"
)

func ExampleRandomString() {
	s := RandomString(20, Digits)
	fmt.Println(s)
}

func ExampleRandomString_default() {
	s := RandomString(20)
	fmt.Println(s)
}
