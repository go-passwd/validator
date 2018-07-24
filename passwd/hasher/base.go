package hasher

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// DefaultSaltLength is a default salt length when is not set manually
	DefaultSaltLength = 20

	// DefaultIter is a default iterations counter when is not set manually
	DefaultIter = 2048
)

// Hasher interface
type Hasher interface {
	// Code returns internal hasher code
	Code() string

	Hash(string) string

	SetPassword(string)

	Check(string) (bool, error)

	String() string
}

// NewFromString returns a new Hasher object who is based on string representation of a hasher (e.x. from database)
func NewFromString(password string) (Hasher, error) {
	p := strings.Split(password, ":")
	switch p[0] {
	case PlainHasher{}.Code():
		return &PlainHasher{Password: &p[1]}, nil
	case MD5Hasher{}.Code():
		iter, _ := strconv.Atoi(p[1])
		return &MD5Hasher{Salt: &p[2], Iter: &iter, Password: &p[3]}, nil
	case SHA224Hasher{}.Code():
		iter, _ := strconv.Atoi(p[1])
		return &SHA224Hasher{Salt: &p[2], Iter: &iter, Password: &p[3]}, nil
	case SHA256Hasher{}.Code():
		iter, _ := strconv.Atoi(p[1])
		return &SHA256Hasher{Salt: &p[2], Iter: &iter, Password: &p[3]}, nil
	case SHA384Hasher{}.Code():
		iter, _ := strconv.Atoi(p[1])
		return &SHA384Hasher{Salt: &p[2], Iter: &iter, Password: &p[3]}, nil
	case SHA512Hasher{}.Code():
		iter, _ := strconv.Atoi(p[1])
		return &SHA512Hasher{Salt: &p[2], Iter: &iter, Password: &p[3]}, nil
	case SHA512_224Hasher{}.Code():
		iter, _ := strconv.Atoi(p[1])
		return &SHA512_224Hasher{Salt: &p[2], Iter: &iter, Password: &p[3]}, nil
	case SHA512_256Hasher{}.Code():
		iter, _ := strconv.Atoi(p[1])
		return &SHA512_256Hasher{Salt: &p[2], Iter: &iter, Password: &p[3]}, nil
	}
	return nil, fmt.Errorf("Unsupported hasher %s", p[0])
}
