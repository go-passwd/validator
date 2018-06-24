package hasher

import (
	"testing"
)

func TestSHA512Hasher_String(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "password"
	h := SHA512Hasher{Salt: &salt, Iter: &iter, Password: &password}

	w := "sha512:1:salt:password"
	g := h.String()
	if g != w {
		t.Errorf("Wanted %s got %s", w, g)
	}
}

func TestSHA512Hasher_Check(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "498931cc17e7391c8e85769bd7b3e91d575ea5c2656eaa16ef6ebda339b21e6ae8a166bd3cb260121f2ae1d8b6d6930970325a1310eeb4c800ef182404408c94"
	h := SHA512Hasher{Salt: &salt, Iter: &iter, Password: &password}

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

func TestSHA512Hasher_Hash(t *testing.T) {
	salt := "salt"
	iter := 1
	password := "498931cc17e7391c8e85769bd7b3e91d575ea5c2656eaa16ef6ebda339b21e6ae8a166bd3cb260121f2ae1d8b6d6930970325a1310eeb4c800ef182404408c94"
	h := SHA512Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	if g != password {
		t.Errorf("Wanted %s got %s", password, g)
	}
}
