# Password validator library for Go

[![Latest tag](https://img.shields.io/github/tag/go-passwd/validator.svg)](https://github.com/go-passwd/validator/tags)
[![GoDoc](https://godoc.org/github.com/go-passwd/validator?status.svg)](https://pkg.go.dev/github.com/go-passwd/validator)
[![Tests status](https://github.com/go-passwd/validator/actions/workflows/test.yml/badge.svg)](https://github.com/go-passwd/validator/actions/workflows/test.yml)
[![Go report](https://goreportcard.com/badge/github.com/go-passwd/validator)](https://goreportcard.com/report/github.com/go-passwd/validator)
[![License](https://img.shields.io/github/license/go-passwd/validator)](./LICENSE)

## Installation

~~~sh
go get -u github.com/go-passwd/validator
~~~

## Usage

~~~go
import "github.com/go-passwd/validator"

passwordValidator := validator.New(validator.MinLength(5, nil), validator.MaxLength(10, nil))
err := passwordValidator.Validate(form.Password)
if err != nil {
  panic(err)
}
~~~

You can pass to every validator functions ``customError`` parameter witch will be returned on error instead of default error.

~~~go
import "github.com/go-passwd/validator"

passwordValidator := validator.New(validator.MinLength(5, errors.New("too short")), validator.MaxLength(10, errors.New("too long")))
err := passwordValidator.Validate(form.Password)
if err != nil {
  panic(err)
}
~~~

## Validators

### CommonPassword

Check if password is a common password.

Common password list is based on list created by Mark Burnett: https://xato.net/passwords/more-top-worst-passwords/

~~~go
passwordValidator := validator.New(validator.CommonPassword(nil))
~~~

### ContainsAtLeast

Count occurrences of a chars and compares it with required value.

~~~go
passwordValidator := validator.New(validator.ContainsAtLeast(5, "abcdefghijklmnopqrstuvwxyz", nil))
~~~

### ContainsOnly

Check if password contains only selected chars.

~~~go
passwordValidator := validator.New(validator.ContainsOnly("abcdefghijklmnopqrstuvwxyz", nil))
~~~

### MaxLength

Check if password length is not greater that defined length.

~~~go
passwordValidator := validator.New(validator.MaxLength(10, nil))
~~~

### MinLength

Check if password length is not lower that defined length.

~~~go
passwordValidator := validator.New(validator.MinLength(5, nil))
~~~

### Noop

Always return custom error.

~~~go
passwordValidator := validator.New(validator.Noop(nil))
~~~

### Regex

Check if password match regexp pattern.

~~~go
passwordValidator := validator.New(validator.Regex("^\\w+$", nil))
~~~

### Similarity

Check if password is sufficiently different from the attributes.

Attributes can be: user login, email, first name, last name, …

~~~go
passwordValidator := validator.New(validator.Similarity([]string{"username", "username@example.com"}, nil, nil))
~~~

### StartsWith

Check if password starts with one of letter.

~~~go
passwordValidator := validator.New(validator.StartsWith("abcdefghijklmnopqrstuvwxyz", nil))
~~~

### Unique

Check if password contains only unique chars.

~~~go
passwordValidator := validator.New(validator.Unique(nil))
~~~
