package hasher

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/tomi77/go-passwd/util"
)

// SHA512Hasher hash password in SHA512
type SHA512Hasher struct {
	Salt     *string
	Iter     *int
	Password *string
}

// Hash a password
func (h SHA512Hasher) Hash(password string) string {
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
		s := sha512.New()
		s.Write(bPassword)
		bPassword = s.Sum(nil)
	}

	return hex.EncodeToString(bPassword)
}

// Check if hashed password is equal stored password hash
func (h *SHA512Hasher) Check(plain string) (bool, error) {
	return h.Hash(plain) == *h.Password, nil
}

func (h *SHA512Hasher) String() string {
	return fmt.Sprintf("sha512:%d:%s:%s", *h.Iter, *h.Salt, *h.Password)
}
