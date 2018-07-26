package hasher

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var password = "plain$password"

func ExampleNewFromString() {
	_, err := NewFromString(password)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestNewFromString_Plain(t *testing.T) {
	g, e := NewFromString("plain$pass")
	assert.Nil(t, e)
	assert.IsType(t, &PlainHasher{}, g)
}

func TestNewFromString_MD5(t *testing.T) {
	g, e := NewFromString("md5$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &MD5Hasher{}, g)
}

func TestNewFromString_SHA1(t *testing.T) {
	g, e := NewFromString("sha1$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &SHA1Hasher{}, g)
}

func TestNewFromString_SHA224(t *testing.T) {
	g, e := NewFromString("sha224$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &SHA224Hasher{}, g)
}

func TestNewFromString_SHA256(t *testing.T) {
	g, e := NewFromString("sha256$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &SHA256Hasher{}, g)
}

func TestNewFromString_SHA384(t *testing.T) {
	g, e := NewFromString("sha384$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &SHA384Hasher{}, g)
}

func TestNewFromString_SHA512(t *testing.T) {
	g, e := NewFromString("sha512$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &SHA512Hasher{}, g)
}

func TestNewFromString_SHA512_224(t *testing.T) {
	g, e := NewFromString("sha512_224$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &SHA512_224Hasher{}, g)
}

func TestNewFromString_SHA512_256(t *testing.T) {
	g, e := NewFromString("sha512_256$1$salt$pass")
	assert.Nil(t, e)
	assert.IsType(t, &SHA512_256Hasher{}, g)
}

func TestNewFromString_bad_hasher(t *testing.T) {
	g, e := NewFromString("qaz123$1")
	assert.NotNil(t, e)
	assert.Nil(t, g)
}
