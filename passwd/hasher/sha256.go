package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/tomi77/go-passwd/passwd/util"
)

// SHA256Hasher hash password in SHA256
type SHA256Hasher struct {
	Salt     *string
	Iter     *int
	Password *string
}

// Code returns internal SHA-224 hasher code
func (h SHA256Hasher) Code() string {
	return "sha256"
}

// Hash a password
func (h SHA256Hasher) Hash(password string) string {
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
		s := sha256.New()
		s.Write(bPassword)
		bPassword = s.Sum(nil)
	}

	return hex.EncodeToString(bPassword)
}

// SetPassword sets a password
func (h *SHA256Hasher) SetPassword(plain string) {
	hash := h.Hash(plain)
	h.Password = &hash
}

// Check if hashed password is equal stored password hash
func (h *SHA256Hasher) Check(plain string) (bool, error) {
	return h.Hash(plain) == *h.Password, nil
}

func (h *SHA256Hasher) String() string {
	return fmt.Sprintf("%s:%d:%s:%s", h.Code(), *h.Iter, *h.Salt, *h.Password)
}
