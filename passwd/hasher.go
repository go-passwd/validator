package passwd

import "github.com/tomi77/go-passwd/passwd/hasher"

// NewSHA512Hasher returns new default SHA512 hasher
func NewSHA512Hasher() hasher.SHA512Hasher {
	return hasher.SHA512Hasher{}
}

// NewPlainHasher returns new default plain hasher
func NewPlainHasher() hasher.PlainHasher {
	return hasher.PlainHasher{}
}

// NewMD5Hasher returns new default MD5 hasher
func NewMD5Hasher() hasher.MD5Hasher {
	return hasher.MD5Hasher{}
}
