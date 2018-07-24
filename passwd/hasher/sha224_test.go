package hasher

import (
	"testing"
)

func TestSHA224Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA224Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha224:1:salt:password"
	g := h.String()
	if g != w {
		t.Errorf("Wanted %s got %s", w, g)
	}
}

func TestSHA224Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "f64671af1dd46e4a00a48a2c7c6a3658d107507391b6eb0d9111b2b3d326512b"
	h := SHA224Hasher{Salt: &salt, Iter: &iter, Password: &password}

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

func TestSHA224Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "f64671af1dd46e4a00a48a2c7c6a3658d107507391b6eb0d9111b2b3d326512b"
	h := SHA224Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	if g != password {
		t.Errorf("Wanted %s got %s", password, g)
	}
}
