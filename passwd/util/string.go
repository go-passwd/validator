package util

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// Base characters allowed in a password
const (
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits       = "1234567890"
)

// RandomString returns a random string of length n consisting of lower letters, upper letters and digitis
func RandomString(n int, baseChars ...string) string {
	var letterBytes string
	letterBytes = strings.Join(baseChars, "")
	if len(letterBytes) == 0 {
		letterBytes = LowerLetters + UpperLetters + Digits
	}
	b := make([]byte, n)
	letterBytesLength := big.NewInt(int64(len(letterBytes)))
	for i := range b {
		idx, _ := rand.Int(rand.Reader, letterBytesLength)
		b[i] = letterBytes[idx.Int64()]
	}

	return string(b)
}
