package hasher

import (
	"testing"
)

func TestSHA256Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA256Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha256$1$salt$password"
	g := h.String()
	if g != w {
		t.Errorf("Wanted %s got %s", w, g)
	}
}

func TestSHA256Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "13601bda4ea78e55a07b98866d2be6be0744e3866f13c00c811cab608a28f322"
	h := SHA256Hasher{Salt: &salt, Iter: &iter, Password: &password}

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

func TestSHA256Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "13601bda4ea78e55a07b98866d2be6be0744e3866f13c00c811cab608a28f322"
	h := SHA256Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	if g != password {
		t.Errorf("Wanted %s got %s", password, g)
	}
}
