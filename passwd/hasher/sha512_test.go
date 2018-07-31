package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA512Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA512Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha512$1$salt$password"
	g := h.String()
	assert.Equal(t, w, g)
}

func TestSHA512Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "2908d2c28dfc047741fc590a026ffade237ab2ba7e1266f010fe49bde548b5987a534a86655a0d17f336588e540cd66f67234b152bbb645b4bb85758a1325d64"
	h := SHA512Hasher{Salt: &salt, Iter: &iter, Password: &password}

	check := h.Check("password")
	assert.Truef(t, check, "Passwords are equal")

	check = h.Check("password2")
	assert.Falsef(t, check, "Passwords are not equal")
}

func TestSHA512Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "2908d2c28dfc047741fc590a026ffade237ab2ba7e1266f010fe49bde548b5987a534a86655a0d17f336588e540cd66f67234b152bbb645b4bb85758a1325d64"
	h := SHA512Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	assert.Equal(t, password, g)
}

func TestSHA512Hasher_Hash_Empty(t *testing.T) {
	h := SHA512Hasher{}
	h.Hash("password")
	assert.NotNil(t, h.Iter)
	assert.NotNil(t, h.Salt)
}

func TestSHA512Hasher_SetPassword(t *testing.T) {
	h := SHA512Hasher{}
	h.SetPassword("password")
	assert.NotNil(t, h.Password)
}
