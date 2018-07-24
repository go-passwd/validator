package passwd

import "github.com/tomi77/go-passwd/passwd/hasher"

// Check function compares password in plain text with string representation of a hasher (e.x. from database)
func Check(plain, hashed string) (bool, error) {
	_hasher, err := hasher.NewFromString(hashed)
	if err != nil {
		return false, err
	}
	check, err := _hasher.Check(plain)
	return check, err
}
