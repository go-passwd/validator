package passwd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomi77/go-passwd/passwd/hasher"
)

var strPassword = "plain$password"

func ExampleNewFromString() {
	h, err := NewFromString(strPassword)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(h.Code())
	// Output:
	// plain
}

func TestNewFromString_Plain(t *testing.T) {
	g, e := NewFromString("plain$pass")
	assert.Nil(t, e)
	assert.IsType(t, &hasher.PlainHasher{}, g)
}

func TestNewFromString_MD5(t *testing.T) {
	g, e := NewFromString("md5$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &hasher.MD5Hasher{}, g)
}

func TestNewFromString_SHA1(t *testing.T) {
	g, e := NewFromString("sha1$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &hasher.SHA1Hasher{}, g)
}

func TestNewFromString_SHA224(t *testing.T) {
	g, e := NewFromString("sha224$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &hasher.SHA224Hasher{}, g)
}

func TestNewFromString_SHA256(t *testing.T) {
	g, e := NewFromString("sha256$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &hasher.SHA256Hasher{}, g)
}

func TestNewFromString_SHA384(t *testing.T) {
	g, e := NewFromString("sha384$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &hasher.SHA384Hasher{}, g)
}

func TestNewFromString_SHA512(t *testing.T) {
	g, e := NewFromString("sha512$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &hasher.SHA512Hasher{}, g)
}

func TestNewFromString_SHA512_224(t *testing.T) {
	g, e := NewFromString("sha512_224$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &hasher.SHA512_224Hasher{}, g)
}

func TestNewFromString_SHA512_256(t *testing.T) {
	g, e := NewFromString("sha512_256$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &hasher.SHA512_256Hasher{}, g)
}

func TestNewFromString_bad_hasher(t *testing.T) {
	g, e := NewFromString("qaz123$1")
	assert.NotNil(t, e)
	assert.Nil(t, g)
}

func TestNewSHA512_256Hasher(t *testing.T) {
	h := NewSHA512_256Hasher()
	assert.Implements(t, (*Hasher)(nil), h)
	assert.IsType(t, &hasher.SHA512_256Hasher{}, h)
}

func TestNewSHA512_224Hasher(t *testing.T) {
	h := NewSHA512_224Hasher()
	assert.Implements(t, (*Hasher)(nil), h)
	assert.IsType(t, &hasher.SHA512_224Hasher{}, h)
}

func TestNewSHA512Hasher(t *testing.T) {
	h := NewSHA512Hasher()
	assert.Implements(t, (*Hasher)(nil), h)
	assert.IsType(t, &hasher.SHA512Hasher{}, h)
}

func TestNewSHA384Hasher(t *testing.T) {
	h := NewSHA384Hasher()
	assert.Implements(t, (*Hasher)(nil), h)
	assert.IsType(t, &hasher.SHA384Hasher{}, h)
}

func TestNewSHA256Hasher(t *testing.T) {
	h := NewSHA256Hasher()
	assert.Implements(t, (*Hasher)(nil), h)
	assert.IsType(t, &hasher.SHA256Hasher{}, h)
}

func TestNewSHA224Hasher(t *testing.T) {
	h := NewSHA224Hasher()
	assert.Implements(t, (*Hasher)(nil), h)
	assert.IsType(t, &hasher.SHA224Hasher{}, h)
}

func TestNewSHA1Hasher(t *testing.T) {
	h := NewSHA1Hasher()
	assert.Implements(t, (*Hasher)(nil), h)
	assert.IsType(t, &hasher.SHA1Hasher{}, h)
}

func TestNewMD5Hasher(t *testing.T) {
	h := NewMD5Hasher()
	assert.Implements(t, (*Hasher)(nil), h)
	assert.IsType(t, &hasher.MD5Hasher{}, h)
}

func TestNewPlainHasher(t *testing.T) {
	h := NewPlainHasher()
	assert.Implements(t, (*Hasher)(nil), h)
	assert.IsType(t, &hasher.PlainHasher{}, h)
}
