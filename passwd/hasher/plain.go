package hasher

import (
	"fmt"
)

// PlainHasher stored password as plain text
type PlainHasher struct {
	Password *string
}

// Code returns internal plain hasher code
func (h PlainHasher) Code() string {
	return "plain"
}

// Hash a password
func (h *PlainHasher) Hash(password string) string {
	return password
}

// SetPassword sets a password
func (h *PlainHasher) SetPassword(password string) {
	_password := h.Hash(password)
	h.Password = &_password
}

// Check if password is equal stored password
func (h *PlainHasher) Check(plain string) bool {
	return plain == *h.Password
}

func (h *PlainHasher) String() string {
	return fmt.Sprintf("%s$%s", h.Code(), *h.Password)
}
