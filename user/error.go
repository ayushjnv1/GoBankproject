package user

import "errors"

var (
	ErrEmptyName     = errors.New("name is not present")
	ErrEmptyEmail    = errors.New("email is not present")
	ErrEmptyRole     = errors.New("role is not present")
	ErrEmptyPassword = errors.New("password not present")
)
