package validator_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-passwd/validator"
)

func TestContainsAtLeast_unicode(t *testing.T) {
	// zestaw znaków z wielobajtowymi rune — musi poprawnie liczyć znaki
	v := validator.ContainsAtLeast("αβγ", 2, nil)
	assert.NoError(t, v("αβ"))
	assert.Error(t, v("α"))
}

func ExampleContainsAtLeast() {
	passwordValidator := validator.ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 4, nil)
	fmt.Println(passwordValidator("password"))
	fmt.Println(passwordValidator("PASSWORD"))
	fmt.Println(passwordValidator("passWORD"))
	// Output:
	// <nil>
	// Password must contain at least 4 chars from abcdefghijklmnopqrstuvwxyz
	// <nil>
}

func ExampleContainsAtLeast_customError() {
	err := errors.New("custom error message")
	passwordValidator := validator.ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 4, err)
	fmt.Println(passwordValidator("PASSWORD"))
	// Output:
	// custom error message
}
