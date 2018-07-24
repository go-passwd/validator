package passwd

import "github.com/tomi77/go-passwd/passwd/hasher"

// NewSHA512_256Hasher returns new default SHA-512/256 hasher
func NewSHA512_256Hasher() hasher.Hasher {
	return &hasher.SHA512_256Hasher{}
}

// NewSHA512_224Hasher returns new default SHA-512/224 hasher
func NewSHA512_224Hasher() hasher.Hasher {
	return &hasher.SHA512_224Hasher{}
}

// NewSHA512Hasher returns new default SHA-512 hasher
func NewSHA512Hasher() hasher.Hasher {
	return &hasher.SHA512Hasher{}
}

// NewSHA384Hasher returns new default SHA-384 hasher
func NewSHA384Hasher() hasher.Hasher {
	return &hasher.SHA384Hasher{}
}

// NewSHA256Hasher returns new default SHA-256 hasher
func NewSHA256Hasher() hasher.Hasher {
	return &hasher.SHA256Hasher{}
}

// NewSHA224Hasher returns new default SHA-224 hasher
func NewSHA224Hasher() hasher.Hasher {
	return &hasher.SHA224Hasher{}
}

// NewSHA1Hasher returns new default SHA-224 hasher
func NewSHA1Hasher() hasher.Hasher {
	return &hasher.SHA1Hasher{}
}

// NewMD5Hasher returns new default MD5 hasher
func NewMD5Hasher() hasher.Hasher {
	return &hasher.MD5Hasher{}
}

// NewPlainHasher returns new default plain hasher
func NewPlainHasher() hasher.Hasher {
	return &hasher.PlainHasher{}
}
