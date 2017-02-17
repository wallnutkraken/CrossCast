package main

import "errors"

var (
	ErrInvalidLogin = errors.New("Invalid login. Remember, everything is case-sensitive.")
	ErrUserAlreadyExists = errors.New("User with that username already exists.")
)
