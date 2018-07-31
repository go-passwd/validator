package passwd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomi77/go-passwd/passwd/validator"
)

func TestValidator_empty(t *testing.T) {
	passwordValidator := Validator{}
	err := passwordValidator.Validate("password")
	assert.NotNil(t, err)
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
	// <nil>
	// <nil>
	// <nil>
	// Password length must be not lower that 5 chars
	// Password length must be not greater that 10 chars
}
