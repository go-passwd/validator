package hasher

import "testing"

func TestPlainHasher_String(t *testing.T) {
	password := "password"
	h := PlainHasher{&password}

	w := "plain:password"
	g := h.String()
	if g != w {
		t.Errorf("Wanted %s got %s", w, g)
	}
}

func TestPlainHasher_Check(t *testing.T) {
	password := "password"
	h := PlainHasher{&password}

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

func TestPlainHasher_Hash(t *testing.T) {
	h := PlainHasher{}

	g := h.Hash("password")
	w := "password"
	if g != w {
		t.Errorf("Wanted %s got %s", w, g)
	}
}
