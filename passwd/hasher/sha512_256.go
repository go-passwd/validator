package hasher

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/tomi77/go-stringgen/randomstring"
)

// SHA512_256Hasher hash password in SHA-512/256
type SHA512_256Hasher struct {
	Salt     *string
	Iter     *int
	Password *string
}

// Code returns internal SHA-512/256 hasher code
func (h SHA512_256Hasher) Code() string {
	return "sha512_256"
}

// Hash a password
func (h *SHA512_256Hasher) Hash(password string) string {
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
		s := sha512.New512_256()
		s.Write(bPassword)
		bPassword = s.Sum(nil)
	}

	return hex.EncodeToString(bPassword)
}

// SetPassword sets a password
func (h *SHA512_256Hasher) SetPassword(plain string) Hasher {
	hash := h.Hash(plain)
	h.Password = &hash
	return h
}

// Check if hashed password is equal stored password hash
func (h *SHA512_256Hasher) Check(plain string) (bool, error) {
	return h.Hash(plain) == *h.Password, nil
}

func (h *SHA512_256Hasher) String() string {
	return fmt.Sprintf("%s$%d$%s$%s", h.Code(), *h.Iter, *h.Salt, *h.Password)
}
