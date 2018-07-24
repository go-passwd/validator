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
	Hash(string) string

	SetPassword(string)

	Check(string) (bool, error)

	String() string
}

// NewFromString returns a new Hasher object who is based on string representation of a hasher (e.x. from database)
func NewFromString(password string) (Hasher, error) {
	p := strings.Split(password, ":")
	switch p[0] {
	case "plain":
		return &PlainHasher{Password: &p[1]}, nil
	case "sha512":
		iter, _ := strconv.Atoi(p[1])
		return &SHA512Hasher{Salt: &p[2], Iter: &iter, Password: &p[3]}, nil
	}
	return nil, fmt.Errorf("Unsupported hasher %s", p[0])
}
