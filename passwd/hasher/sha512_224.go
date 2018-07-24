package hasher

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/tomi77/go-passwd/passwd/util"
)

const _SHA512_224MarshalerCode = "sha512_224"

// SHA512_224Hasher hash password in SHA-512/224
type SHA512_224Hasher struct {
	Salt     *string
	Iter     *int
	Password *string
}

// Hash a password
func (h SHA512_224Hasher) Hash(password string) string {
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
		s := sha512.New512_224()
		s.Write(bPassword)
		bPassword = s.Sum(nil)
	}

	return hex.EncodeToString(bPassword)
}

// SetPassword sets a password
func (h *SHA512_224Hasher) SetPassword(plain string) {
	hash := h.Hash(plain)
	h.Password = &hash
}

// Check if hashed password is equal stored password hash
func (h *SHA512_224Hasher) Check(plain string) (bool, error) {
	return h.Hash(plain) == *h.Password, nil
}

func (h *SHA512_224Hasher) String() string {
	return fmt.Sprintf("%s:%d:%s:%s", _SHA512_224MarshalerCode, *h.Iter, *h.Salt, *h.Password)
}
