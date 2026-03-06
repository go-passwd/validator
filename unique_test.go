package validator_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-passwd/validator"
)

func TestUnique_nonASCIIWithoutDuplicates(t *testing.T) {
	v := validator.Unique(nil)
	assert.NoError(t, v("αβγδ"))
}

func TestUnique_nonASCIIWithDuplicates(t *testing.T) {
	v := validator.Unique(nil)
	assert.Error(t, v("αβγα"))
}

func ExampleUnique() {
	passwordValidator := validator.Unique(nil)
	fmt.Println(passwordValidator("bui87j"))
	fmt.Println(passwordValidator("qwerte"))
	// Output:
	// <nil>
	// the password must contain unique chars
}

func ExampleUnique_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.Unique(err)
	fmt.Println(passwordValidator("bui87j"))
	fmt.Println(passwordValidator("qwerte"))
	// Output:
	// <nil>
	// custom error message
}
