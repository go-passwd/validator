package hasher

import (
	"testing"
)

func TestMD5Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := MD5Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "md5$1$salt$password"
	g := h.String()
	if g != w {
		t.Errorf("Wanted %s got %s", w, g)
	}
}

func TestMD5Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "67a1e09bb1f83f5007dc119c14d663aa"
	h := MD5Hasher{Salt: &salt, Iter: &iter, Password: &password}

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

func TestMD5Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "67a1e09bb1f83f5007dc119c14d663aa"
	h := MD5Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	if g != password {
		t.Errorf("Wanted %s got %s", password, g)
	}
}
