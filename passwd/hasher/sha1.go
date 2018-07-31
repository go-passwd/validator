package hasher

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/tomi77/go-stringgen/randomstring"
)

// SHA1Hasher hash password in SHA-1
type SHA1Hasher struct {
	Salt     *string
	Iter     *int
	Password *string
}

// Code returns internal SHA-224 hasher code
func (h SHA1Hasher) Code() string {
	return "sha1"
}

// Hash a password
func (h *SHA1Hasher) Hash(password string) string {
	if h.Salt == nil {
		salt := randomstring.Generate(DefaultSaltLength)
		h.Salt = &salt
	}

	if h.Iter == nil {
		iter := DefaultIter
		h.Iter = &iter
	}

	bPassword := []byte(*h.Salt + password)
	for i := 0; i < *h.Iter; i++ {
		s := sha1.New()
		s.Write(bPassword)
		bPassword = s.Sum(nil)
	}

	return hex.EncodeToString(bPassword)
}

// SetPassword sets a password
func (h *SHA1Hasher) SetPassword(plain string) {
	hash := h.Hash(plain)
	h.Password = &hash
}

// Check if hashed password is equal stored password hash
func (h *SHA1Hasher) Check(plain string) bool {
	return h.Hash(plain) == *h.Password
}

func (h *SHA1Hasher) String() string {
	return fmt.Sprintf("%s$%d$%s$%s", h.Code(), *h.Iter, *h.Salt, *h.Password)
}
