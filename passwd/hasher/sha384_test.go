package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA384Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA384Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha384$1$salt$password"
	g := h.String()
	assert.Equal(t, w, g)
}

func TestSHA384Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "f4bdac9860c0ceea69fb29efbce24addca5cf1f808925d9433b668528290d5d2c9080f32342175b5124895684db8ba4f"
	h := SHA384Hasher{Salt: &salt, Iter: &iter, Password: &password}

	check := h.Check("password")
	assert.Truef(t, check, "Passwords are equal")

	check = h.Check("password2")
	assert.Falsef(t, check, "Passwords are not equal")
}

func TestSHA384Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "f4bdac9860c0ceea69fb29efbce24addca5cf1f808925d9433b668528290d5d2c9080f32342175b5124895684db8ba4f"
	h := SHA384Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	assert.Equal(t, password, g)
}

func TestSHA384Hasher_Hash_Empty(t *testing.T) {
	h := SHA384Hasher{}
	h.Hash("password")
	assert.NotNil(t, h.Iter)
	assert.NotNil(t, h.Salt)
}

func TestSHA384Hasher_SetPassword(t *testing.T) {
	h := SHA384Hasher{}
	h.SetPassword("password")
	assert.NotNil(t, h.Password)
}
