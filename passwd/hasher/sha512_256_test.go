package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA512_256Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA512_256Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha512_256$1$salt$password"
	g := h.String()
	assert.Equal(t, w, g)
}

func TestSHA512_256Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "0542188208714e74c77107558168d7fea61d85bbab7a5c88af3cf6fa2068b601"
	h := SHA512_256Hasher{Salt: &salt, Iter: &iter, Password: &password}

	check := h.Check("password")
	assert.Truef(t, check, "Passwords are equal")

	check = h.Check("password2")
	assert.Falsef(t, check, "Passwords are not equal")
}

func TestSHA512_256Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "0542188208714e74c77107558168d7fea61d85bbab7a5c88af3cf6fa2068b601"
	h := SHA512_256Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	assert.Equal(t, password, g)
}

func TestSHA512_256Hasher_Hash_Empty(t *testing.T) {
	h := SHA512_256Hasher{}
	h.Hash("password")
	assert.NotNil(t, h.Iter)
	assert.NotNil(t, h.Salt)
}

func TestSHA512_256Hasher_SetPassword(t *testing.T) {
	h := SHA512_256Hasher{}
	h.SetPassword("password")
	assert.NotNil(t, h.Password)
}
