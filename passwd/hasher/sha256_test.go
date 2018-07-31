package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA256Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA256Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha256$1$salt$password"
	g := h.String()
	assert.Equal(t, w, g)
}

func TestSHA256Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "13601bda4ea78e55a07b98866d2be6be0744e3866f13c00c811cab608a28f322"
	h := SHA256Hasher{Salt: &salt, Iter: &iter, Password: &password}

	check := h.Check("password")
	assert.Truef(t, check, "Passwords are equal")

	check = h.Check("password2")
	assert.Falsef(t, check, "Passwords are not equal")
}

func TestSHA256Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "13601bda4ea78e55a07b98866d2be6be0744e3866f13c00c811cab608a28f322"
	h := SHA256Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	assert.Equal(t, password, g)
}

func TestSHA256Hasher_Hash_Empty(t *testing.T) {
	h := SHA256Hasher{}
	h.Hash("password")
	assert.NotNil(t, h.Iter)
	assert.NotNil(t, h.Salt)
}

func TestSHA256Hasher_SetPassword(t *testing.T) {
	h := SHA256Hasher{}
	h.SetPassword("password")
	assert.NotNil(t, h.Password)
}
