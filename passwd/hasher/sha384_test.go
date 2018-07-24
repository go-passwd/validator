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
	password := "574b75d17fa4fb13617b93d7e6be3914e9cf068ca0a6573a947c616481b3dec31f0f1e5a133f4ec5d365b921fdc0a9c0"
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
	password := "574b75d17fa4fb13617b93d7e6be3914e9cf068ca0a6573a947c616481b3dec31f0f1e5a133f4ec5d365b921fdc0a9c0"
	h := SHA384Hasher{Salt: &salt, Iter: &iter}

	g := h.Hash("password")
	if g != password {
		t.Errorf("Wanted %s got %s", password, g)
	}
}
