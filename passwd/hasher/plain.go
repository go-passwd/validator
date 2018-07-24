package hasher

import (
	"fmt"
)

// PlainHasher stored password in plain text
type PlainHasher struct {
	Password *string
}

// Hash a password
func (h PlainHasher) Hash(password string) string {
	return password
}

// SetPassword sets a password
func (h *PlainHasher) SetPassword(password string) {
	_password := h.Hash(password)
	h.Password = &_password
}

// Check if password is equal stored password
func (h *PlainHasher) Check(plain string) (bool, error) {
	return plain == *h.Password, nil
}

func (h *PlainHasher) String() string {
	return fmt.Sprintf("plain:%s", *h.Password)
}
