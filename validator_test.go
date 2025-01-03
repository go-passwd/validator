package validator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-passwd/validator"
)

func TestValidator_empty(t *testing.T) {
	passwordValidator := validator.Validator{}
	err := passwordValidator.Validate("password")
	require.Error(t, err)
}

func TestNewValidator(t *testing.T) {
	passwordValidator := validator.New(validator.MinLength(5, nil), validator.MaxLength(10, nil))
	require.Len(t, *passwordValidator, 2)
	require.NoError(t, passwordValidator.Validate("password"))
	require.NoError(t, passwordValidator.Validate("pass1"))
	require.NoError(t, passwordValidator.Validate("password12"))
	require.Error(t, passwordValidator.Validate("pass"))
	require.Error(t, passwordValidator.Validate("password123"))
}

func ExampleValidator_Validate() {
	passwordValidator := validator.Validator{validator.MinLength(5, nil), validator.MaxLength(10, nil)}
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
