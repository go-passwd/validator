package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA1Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA1Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha1$1$salt$password"
	g := h.String()
	assert.Equal(t, w, g)
}

func TestSHA1Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "59b3e8d637cf97edbe2384cf59cb7453dfe30789"
	h := SHA1Hasher{Salt: &salt, Iter: &iter, Password: &password}

	check := h.Check("password")
	assert.Truef(t, check, "Passwords are equal")

	check = h.Check("password2")
	assert.Falsef(t, check, "Passwords are not equal")
}

func TestSHA1Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "59b3e8d637cf97edbe2384cf59cb7453dfe30789"
	h := SHA1Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	assert.Equal(t, password, g)
}

func TestSHA1Hasher_Hash_Empty(t *testing.T) {
	h := SHA1Hasher{}
	h.Hash("password")
	assert.NotNil(t, h.Iter)
	assert.NotNil(t, h.Salt)
}

func TestSHA1Hasher_SetPassword(t *testing.T) {
	h := SHA1Hasher{}
	h.SetPassword("password")
	assert.NotNil(t, h.Password)
}
