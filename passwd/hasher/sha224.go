package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/tomi77/go-passwd/passwd/util"
)

// SHA224Hasher hash password in SHA-224
type SHA224Hasher struct {
	Salt     *string
	Iter     *int
	Password *string
}

// Code returns internal SHA-224 hasher code
func (h SHA224Hasher) Code() string {
	return "sha224"
}

// Hash a password
func (h SHA224Hasher) Hash(password string) string {
	if h.Salt == nil {
		salt := util.RandomString(DefaultSaltLength)
		h.Salt = &salt
	}

	if h.Iter == nil {
		iter := DefaultIter
		h.Iter = &iter
	}

	bPassword := []byte(password + ":" + *h.Salt)
	for i := 0; i < *h.Iter; i++ {
		s := sha256.New224()
		s.Write(bPassword)
		bPassword = s.Sum(nil)
	}

	return hex.EncodeToString(bPassword)
}

// SetPassword sets a password
func (h *SHA224Hasher) SetPassword(plain string) Hasher {
	hash := h.Hash(plain)
	h.Password = &hash
	return h
}

// Check if hashed password is equal stored password hash
func (h *SHA224Hasher) Check(plain string) (bool, error) {
	return h.Hash(plain) == *h.Password, nil
}

func (h *SHA224Hasher) String() string {
	return fmt.Sprintf("%s$%d$%s$%s", h.Code(), *h.Iter, *h.Salt, *h.Password)
}
