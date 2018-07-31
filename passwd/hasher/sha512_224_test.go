package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA512_224Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA512_224Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha512_224$1$salt$password"
	g := h.String()
	assert.Equal(t, w, g)
}

func TestSHA512_224Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "8a63ceaac7f7b60975d61bc1b8b1c76ec7de6c0226b7af60633ec50e"
	h := SHA512_224Hasher{Salt: &salt, Iter: &iter, Password: &password}

	check := h.Check("password")
	assert.Truef(t, check, "Passwords are equal")

	check = h.Check("password2")
	assert.Falsef(t, check, "Passwords are not equal")
}

func TestSHA512_224Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "8a63ceaac7f7b60975d61bc1b8b1c76ec7de6c0226b7af60633ec50e"
	h := SHA512_224Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	assert.Equal(t, password, g)
}

func TestSHA512_224Hasher_Hash_Empty(t *testing.T) {
	h := SHA512_224Hasher{}
	h.Hash("password")
	assert.NotNil(t, h.Iter)
	assert.NotNil(t, h.Salt)
}

func TestSHA512_224Hasher_SetPassword(t *testing.T) {
	h := SHA512_224Hasher{}
	h.SetPassword("password")
	assert.NotNil(t, h.Password)
}
