package validator_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-passwd/validator"
)

func TestCommonPassword_caseInsensitive(t *testing.T) {
	v := validator.CommonPassword(nil)
	// "password" is in the list — all case variants must be rejected
	assert.Error(t, v("password"))
	assert.Error(t, v("Password"))
	assert.Error(t, v("PASSWORD"))
	assert.Error(t, v("pAsSwOrD"))
}

func ExampleCommonPassword() {
	passwordValidator := validator.CommonPassword(nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("qaz123"))
	fmt.Println(passwordValidator("pa$$w0rd@"))
	fmt.Println(passwordValidator("Mod7tygrysow"))
	// Output:
	// Password can't be a commonly used password
	// Password can't be a commonly used password
	// <nil>
	// Password can't be a commonly used password
}

func ExampleCommonPassword_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.CommonPassword(err)
	fmt.Println(passwordValidator("password"))
	// Output:
	// custom error message
}
