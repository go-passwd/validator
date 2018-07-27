package util

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	s := RandomString(100, Digits, LowerLetters)
	assert.Len(t, s, 100)
	for _, l := range strings.Split(UpperLetters, "") {
		assert.NotContains(t, s, l)
	}

	s = RandomString(20, Digits)
	assert.Len(t, s, 20)
	for _, l := range strings.Split(UpperLetters+LowerLetters, "") {
		assert.NotContains(t, s, l)
	}
}

func TestRandomString_Default(t *testing.T) {
	s := RandomString(30)
	assert.Len(t, s, 30)
}

func ExampleRandomString() {
	s := RandomString(20, Digits)
	fmt.Println(s)
}

func ExampleRandomString_default() {
	s := RandomString(20)
	fmt.Println(s)
}
