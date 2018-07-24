package hasher

import (
	"testing"
)

func TestSHA1Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA1Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha1$1$salt$password"
	g := h.String()
	if g != w {
		t.Errorf("Wanted %s got %s", w, g)
	}
}

func TestSHA1Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "676f03a8c8530384eb7551b1685f2828546625d9"
	h := SHA1Hasher{Salt: &salt, Iter: &iter, Password: &password}

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

func TestSHA1Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "676f03a8c8530384eb7551b1685f2828546625d9"
	h := SHA1Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	if g != password {
		t.Errorf("Wanted %s got %s", password, g)
	}
}
