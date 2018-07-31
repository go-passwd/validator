package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := MD5Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "md5$1$salt$password"
	g := h.String()
	assert.Equal(t, w, g)
}

func TestMD5Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "67a1e09bb1f83f5007dc119c14d663aa"
	h := MD5Hasher{Salt: &salt, Iter: &iter, Password: &password}

	check := h.Check("password")
	assert.Truef(t, check, "Passwords are equal")

	check = h.Check("password2")
	assert.Falsef(t, check, "Passwords are not equal")
}

func TestMD5Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "67a1e09bb1f83f5007dc119c14d663aa"
	h := MD5Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	assert.Equal(t, password, g)
}

func TestMD5Hasher_Hash_Empty(t *testing.T) {
	h := MD5Hasher{}
	h.Hash("password")
	assert.NotNil(t, h.Iter)
	assert.NotNil(t, h.Salt)
}

func TestMD5Hasher_SetPassword(t *testing.T) {
	h := MD5Hasher{}
	h.SetPassword("password")
	assert.NotNil(t, h.Password)
}
