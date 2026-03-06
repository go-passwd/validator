package validator_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-passwd/validator"
)

func TestMaxLength_unicode(t *testing.T) {
	v := validator.MaxLength(4, nil)
	assert.NoError(t, v("łódź")) // 4 runy, 7 bajtów — powinno przejść
	assert.Error(t, v("łódźx")) // 5 run — powinno nie przejść
}

func ExampleMaxLength() {
	passwordValidator := validator.MaxLength(5, nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("pass"))
	fmt.Println(passwordValidator("passw"))
	// Output:
	// Password length must be not greater that 5 chars
	// <nil>
	// <nil>
}

func ExampleMaxLength_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.MaxLength(5, err)
	fmt.Println(passwordValidator("password"))
	// Output:
	// custom error message
}
