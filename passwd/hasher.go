package passwd

import "github.com/tomi77/go-passwd/passwd/hasher"

// NewSHA512_224Hasher returns new default SHA-512/224 hasher
func NewSHA512_224Hasher() hasher.SHA512_224Hasher {
	return hasher.SHA512_224Hasher{}
}

// NewSHA512Hasher returns new default SHA-512 hasher
func NewSHA512Hasher() hasher.SHA512Hasher {
	return hasher.SHA512Hasher{}
}

// NewSHA384Hasher returns new default SHA-384 hasher
func NewSHA384Hasher() hasher.SHA384Hasher {
	return hasher.SHA384Hasher{}
}

// NewPlainHasher returns new default plain hasher
func NewPlainHasher() hasher.PlainHasher {
	return hasher.PlainHasher{}
}

// NewMD5Hasher returns new default MD5 hasher
func NewMD5Hasher() hasher.MD5Hasher {
	return hasher.MD5Hasher{}
}
