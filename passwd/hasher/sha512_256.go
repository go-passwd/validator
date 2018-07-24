package hasher

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/tomi77/go-passwd/passwd/util"
)

const _SHA512_256MarshalerCode = "sha512_256"

// SHA512_256Hasher hash password in SHA-512/256
type SHA512_256Hasher struct {
	Salt     *string
	Iter     *int
	Password *string
}

// Hash a password
func (h SHA512_256Hasher) Hash(password string) string {
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
		s := sha512.New512_256()
		s.Write(bPassword)
		bPassword = s.Sum(nil)
	}

	return hex.EncodeToString(bPassword)
}

// SetPassword sets a password
func (h *SHA512_256Hasher) SetPassword(plain string) {
	hash := h.Hash(plain)
	h.Password = &hash
}

// Check if hashed password is equal stored password hash
func (h *SHA512_256Hasher) Check(plain string) (bool, error) {
	return h.Hash(plain) == *h.Password, nil
}

func (h *SHA512_256Hasher) String() string {
	return fmt.Sprintf("%s:%d:%s:%s", _SHA512_256MarshalerCode, *h.Iter, *h.Salt, *h.Password)
}
