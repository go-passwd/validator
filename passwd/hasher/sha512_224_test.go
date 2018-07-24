package hasher

import (
	"testing"
)

func TestSHA512_224Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA512_224Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha512_224:1:salt:password"
	g := h.String()
	if g != w {
		t.Errorf("Wanted %s got %s", w, g)
	}
}

func TestSHA512_224Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "2ac5d1906aac52d33776fa0d8e6e28922e191a11dbb15834ea3eed90"
	h := SHA512_224Hasher{Salt: &salt, Iter: &iter, Password: &password}

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

func TestSHA512_224Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "2ac5d1906aac52d33776fa0d8e6e28922e191a11dbb15834ea3eed90"
	h := SHA512_224Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	if g != password {
		t.Errorf("Wanted %s got %s", password, g)
	}
}
