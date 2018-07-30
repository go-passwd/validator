package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/tomi77/go-stringgen/randomstring"
)

// MD5Hasher hash password in MD5
type MD5Hasher struct {
	Salt     *string
	Iter     *int
	Password *string
}

// Code returns internal MD5 hasher code
func (h MD5Hasher) Code() string {
	return "md5"
}

// Hash a password
func (h *MD5Hasher) Hash(password string) string {
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
		s := md5.New()
		s.Write(bPassword)
		bPassword = s.Sum(nil)
	}

	return hex.EncodeToString(bPassword)
}

// SetPassword sets a password
func (h *MD5Hasher) SetPassword(plain string) Hasher {
	hash := h.Hash(plain)
	h.Password = &hash
	return h
}

// Check if hashed password is equal stored password hash
func (h *MD5Hasher) Check(plain string) (bool, error) {
	return h.Hash(plain) == *h.Password, nil
}

func (h *MD5Hasher) String() string {
	return fmt.Sprintf("%s$%d$%s$%s", h.Code(), *h.Iter, *h.Salt, *h.Password)
}
