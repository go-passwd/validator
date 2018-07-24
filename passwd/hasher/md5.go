package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/tomi77/go-passwd/passwd/util"
)

const _MD5MarshalerCode = "md5"

// MD5Hasher hash password in MD5
type MD5Hasher struct {
	Salt     *string
	Iter     *int
	Password *string
}

// Hash a password
func (h MD5Hasher) Hash(password string) string {
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
		s := md5.New()
		s.Write(bPassword)
		bPassword = s.Sum(nil)
	}

	return hex.EncodeToString(bPassword)
}

// SetPassword sets a password
func (h *MD5Hasher) SetPassword(plain string) {
	hash := h.Hash(plain)
	h.Password = &hash
}

// Check if hashed password is equal stored password hash
func (h *MD5Hasher) Check(plain string) (bool, error) {
	return h.Hash(plain) == *h.Password, nil
}

func (h *MD5Hasher) String() string {
	return fmt.Sprintf("%s:%d:%s:%s", _MD5MarshalerCode, *h.Iter, *h.Salt, *h.Password)
}
