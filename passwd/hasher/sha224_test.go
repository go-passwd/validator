package hasher

import (
	"testing"
)

func TestSHA224Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA224Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha224$1$salt$password"
	g := h.String()
	if g != w {
		t.Errorf("Wanted %s got %s", w, g)
	}
}

func TestSHA224Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "8be78fd2b899875ec45e23d239c5b671a236484395676a56e74a0466"
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
	password := "8be78fd2b899875ec45e23d239c5b671a236484395676a56e74a0466"
	h := SHA224Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	if g != password {
		t.Errorf("Wanted %s got %s", password, g)
	}
}
