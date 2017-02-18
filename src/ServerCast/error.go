package main

import "errors"

var (
	ErrInvalidLogin = errors.New("Invalid login. Remember, everything is case-sensitive.")
	ErrUserAlreadyExists = errors.New("User with that username already exists.")
	ErrNoSuchUser = errors.New("Invalid access token")
	ErrTokenExpired = errors.New("Token has expired. Please re-log in.")
)
