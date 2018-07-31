package passwd

// Check function compares password in plain text with string representation of a hasher (e.x. from database)
func Check(plain, hashed string) (bool, error) {
	_hasher, err := NewFromString(hashed)
	if err != nil {
		return false, err
	}
	return _hasher.Check(plain), nil
}
