package passwd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomi77/go-passwd/passwd/validator"
)

func TestValidator_empty(t *testing.T) {
	passwordValidator := Validator{}
	assert.False(t, passwordValidator.Validate("password"))
}

var password = "password"

func ExampleValidator() {
	passwordValidator := Validator{validator.MinLength(5), validator.MaxLength(10)}
	passwordValidator.Validate(password)
}

func ExampleValidator_Validate() {
	passwordValidator := Validator{validator.MinLength(5), validator.MaxLength(10)}
	fmt.Println(passwordValidator.Validate("password"))
	fmt.Println(passwordValidator.Validate("pass1"))
	fmt.Println(passwordValidator.Validate("password12"))
	fmt.Println(passwordValidator.Validate("pass"))
	fmt.Println(passwordValidator.Validate("password123"))
	// Output:
	// true
	// true
	// true
	// false
	// false
}
