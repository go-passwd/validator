package validator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_empty(t *testing.T) {
	passwordValidator := Validator{}
	err := passwordValidator.Validate("password")
	assert.NotNil(t, err)
}

var password = "password"

func ExampleNew() {
	New(MinLength(5, nil), MaxLength(10, nil))
}

func TestNewValidator(t *testing.T) {
	passwordValidator := New(MinLength(5, nil), MaxLength(10, nil))
	assert.Len(t, *passwordValidator, 2)
	assert.Nil(t, passwordValidator.Validate("password"))
	assert.Nil(t, passwordValidator.Validate("pass1"))
	assert.Nil(t, passwordValidator.Validate("password12"))
	assert.Error(t, passwordValidator.Validate("pass"))
	assert.Error(t, passwordValidator.Validate("password123"))
}

func ExampleValidator_Validate() {
	passwordValidator := Validator{MinLength(5, nil), MaxLength(10, nil)}
	fmt.Println(passwordValidator.Validate("password"))
	fmt.Println(passwordValidator.Validate("pass1"))
	fmt.Println(passwordValidator.Validate("password12"))
	fmt.Println(passwordValidator.Validate("pass"))
	fmt.Println(passwordValidator.Validate("password123"))
	// Output:
	// <nil>
	// <nil>
	// <nil>
	// Password length must be not lower that 5 chars
	// Password length must be not greater that 10 chars
}
