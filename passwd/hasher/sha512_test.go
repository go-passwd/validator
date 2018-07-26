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

	check, err := h.Check("password")
	assert.Nil(t, err)
	assert.Truef(t, check, "Passwords are equal")

	check, err = h.Check("password2")
	assert.Nil(t, err)
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
