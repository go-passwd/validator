package hasher

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/tomi77/go-passwd/passwd/util"
)

// SHA384Hasher hash password in SHA-384
type SHA384Hasher struct {
	Salt     *string
	Iter     *int
	Password *string
}

// Code returns internal SHA-384 hasher code
func (h SHA384Hasher) Code() string {
	return "sha384"
}

// Hash a password
func (h SHA384Hasher) Hash(password string) string {
	if h.Salt == nil {
		salt := util.RandomString(DefaultSaltLength)
		h.Salt = &salt
	}

	if h.Iter == nil {
		iter := DefaultIter
		h.Iter = &iter
	}

	bPassword := []byte(*h.Salt + password)
	for i := 0; i < *h.Iter; i++ {
		s := sha512.New384()
		s.Write(bPassword)
		bPassword = s.Sum(nil)
	}

	return hex.EncodeToString(bPassword)
}

// SetPassword sets a password
func (h *SHA384Hasher) SetPassword(plain string) Hasher {
	hash := h.Hash(plain)
	h.Password = &hash
	return h
}

// Check if hashed password is equal stored password hash
func (h *SHA384Hasher) Check(plain string) (bool, error) {
	return h.Hash(plain) == *h.Password, nil
}

func (h *SHA384Hasher) String() string {
	return fmt.Sprintf("%s$%d$%s$%s", h.Code(), *h.Iter, *h.Salt, *h.Password)
}
