package passwd

import "github.com/tomi77/go-passwd/passwd/hasher"

// NewSHA512_256Hasher returns new default SHA-512/256 hasher
func NewSHA512_256Hasher() hasher.SHA512_256Hasher {
	return hasher.SHA512_256Hasher{}
}

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

// NewSHA224Hasher returns new default SHA-224 hasher
func NewSHA224Hasher() hasher.SHA224Hasher {
	return hasher.SHA224Hasher{}
}

// NewPlainHasher returns new default plain hasher
func NewPlainHasher() hasher.PlainHasher {
	return hasher.PlainHasher{}
}

// NewMD5Hasher returns new default MD5 hasher
func NewMD5Hasher() hasher.MD5Hasher {
	return hasher.MD5Hasher{}
}
