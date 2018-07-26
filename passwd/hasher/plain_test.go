package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlainHasher_String(t *testing.T) {
	password := "password"
	h := PlainHasher{&password}

	w := "plain$password"
	g := h.String()
	assert.Equal(t, w, g)
}

func TestPlainHasher_Check(t *testing.T) {
	password := "password"
	h := PlainHasher{&password}

	check, err := h.Check("password")
	assert.Nil(t, err)
	assert.Truef(t, check, "Passwords are equal")

	check, err = h.Check("password2")
	assert.Nil(t, err)
	assert.Falsef(t, check, "Passwords are not equal")
}

func TestPlainHasher_Hash(t *testing.T) {
	h := PlainHasher{}

	g := h.Hash("password")
	assert.Equal(t, "password", g)
}
