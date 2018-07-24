package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/tomi77/go-passwd/passwd/util"
)

const _SHA224MarshalerCode = "sha224"

// SHA224Hasher hash password in SHA512
type SHA224Hasher struct {
	Salt     *string
	Iter     *int
	Password *string
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
		s := sha256.New()
		s.Write(bPassword)
		bPassword = s.Sum(nil)
	}

	return hex.EncodeToString(bPassword)
}

// SetPassword sets a password
func (h *SHA224Hasher) SetPassword(plain string) {
	hash := h.Hash(plain)
	h.Password = &hash
}

// Check if hashed password is equal stored password hash
func (h *SHA224Hasher) Check(plain string) (bool, error) {
	return h.Hash(plain) == *h.Password, nil
}

func (h *SHA224Hasher) String() string {
	return fmt.Sprintf("%s:%d:%s:%s", _SHA224MarshalerCode, *h.Iter, *h.Salt, *h.Password)
}
