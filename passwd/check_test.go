package passwd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleCheck() {
	correct, err := Check("password", "plain$password")
	if err != nil {
		panic(err)
	}
	if !correct {
		panic("Incorrect password!")
	}
	fmt.Println("OK")
	// Output:
	// OK
}

func TestCheck_incorrectPassword(t *testing.T) {
	correct, err := Check("password", "plained$password")
	assert.False(t, correct)
	assert.Error(t, err)
}
