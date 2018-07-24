package hasher

import (
	"testing"
)

func TestSHA384Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA384Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha384$1$salt$password"
	g := h.String()
	if g != w {
		t.Errorf("Wanted %s got %s", w, g)
	}
}

func TestSHA384Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "f4bdac9860c0ceea69fb29efbce24addca5cf1f808925d9433b668528290d5d2c9080f32342175b5124895684db8ba4f"
	h := SHA384Hasher{Salt: &salt, Iter: &iter, Password: &password}

	check, err := h.Check("password")
	if err != nil {
		t.Error(err)
	}
	if !check {
		t.Error("Passwords are equal")
	}

	check, err = h.Check("password2")
	if err != nil {
		t.Error(err)
	}
	if check {
		t.Error("Passwords are not equal")
	}
}

func TestSHA384Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "f4bdac9860c0ceea69fb29efbce24addca5cf1f808925d9433b668528290d5d2c9080f32342175b5124895684db8ba4f"
	h := SHA384Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	if g != password {
		t.Errorf("Wanted %s got %s", password, g)
	}
}
