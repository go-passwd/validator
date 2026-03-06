package validator_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-passwd/validator"
)

func TestMinLength_unicode(t *testing.T) {
	v := validator.MinLength(4, nil)
	assert.NoError(t, v("łódź")) // 4 runy, 7 bajtów — powinno przejść
	assert.Error(t, v("łód"))   // 3 runy — powinno nie przejść
}

func ExampleMinLength() {
	passwordValidator := validator.MinLength(5, nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("pass"))
	fmt.Println(passwordValidator("passw"))
	// Output:
	// <nil>
	// Password length must be not lower that 5 chars
	// <nil>
}

func ExampleMinLength_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.MinLength(5, err)
	fmt.Println(passwordValidator("pass"))
	// Output:
	// custom error message
}
