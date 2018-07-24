package hasher

import (
	"testing"
)

func TestSHA512_256Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA512_256Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha512_256$1$salt$password"
	g := h.String()
	if g != w {
		t.Errorf("Wanted %s got %s", w, g)
	}
}

func TestSHA512_256Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "0542188208714e74c77107558168d7fea61d85bbab7a5c88af3cf6fa2068b601"
	h := SHA512_256Hasher{Salt: &salt, Iter: &iter, Password: &password}

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

func TestSHA512_256Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "0542188208714e74c77107558168d7fea61d85bbab7a5c88af3cf6fa2068b601"
	h := SHA512_256Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	if g != password {
		t.Errorf("Wanted %s got %s", password, g)
	}
}
