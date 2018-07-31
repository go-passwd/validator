package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA224Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA224Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha224$1$salt$password"
	g := h.String()
	assert.Equal(t, w, g)
}

func TestSHA224Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "5fe31e9aab92219c047273219ab12eba400c9312ae74258706f144e1"
	h := SHA224Hasher{Salt: &salt, Iter: &iter, Password: &password}

	check := h.Check("password")
	assert.Truef(t, check, "Passwords are equal")

	check = h.Check("password2")
	assert.Falsef(t, check, "Passwords are not equal")
}

func TestSHA224Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "5fe31e9aab92219c047273219ab12eba400c9312ae74258706f144e1"
	h := SHA224Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	assert.Equal(t, password, g)
}

func TestSHA224Hasher_Hash_Empty(t *testing.T) {
	h := SHA224Hasher{}
	h.Hash("password")
	assert.NotNil(t, h.Iter)
	assert.NotNil(t, h.Salt)
}

func TestSHA224Hasher_SetPassword(t *testing.T) {
	h := SHA224Hasher{}
	h.SetPassword("password")
	assert.NotNil(t, h.Password)
}
