# Password validator library for Go

[![Build Status](https://travis-ci.org/go-passwd/validator.svg?branch=master)](https://travis-ci.org/go-passwd/validator)
[![Coverage Status](https://coveralls.io/repos/github/go-passwd/validator/badge.svg?branch=master)](https://coveralls.io/github/go-passwd/validator?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-passwd/validator)](https://goreportcard.com/report/github.com/go-passwd/validator)
[![GoDoc](https://godoc.org/github.com/go-passwd/validator?status.svg)](https://godoc.org/github.com/go-passwd/validator)

## Installation

~~~sh
go get -u github.com/go-passwd/validator
~~~

## Usage

~~~go
import "github.com/go-passwd/validator"

passwordValidator := validator.New(validator.MinLength(5), validator.MaxLength(10))
err := passwordValidator.Validate(form.Password)
if err != nil {
  panic(err)
}
~~~

## Validators

### MinLength

Check if password length is not lower that defined length.

~~~go
passwordValidator := validator.New(validator.MinLength(5))
~~~

### MaxLength

Check if password length is not greater that defined length.

~~~go
passwordValidator := validator.New(validator.MaxLength(10))
~~~

### ContainsAtLeast

Count occurrences of a chars and compares it with required value.

~~~go
passwordValidator := validator.New(validator.ContainsAtLeast(5, "abcdefghijklmnopqrstuvwxyz")
~~~

### CommonPassword

Check if password is a common password.

Common password list is based on list created by Mark Burnett: https://xato.net/passwords/more-top-worst-passwords/

~~~go
passwordValidator := validator.New(validator.CommonPassword()
~~~
