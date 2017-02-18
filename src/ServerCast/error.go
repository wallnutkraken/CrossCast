package main

import "errors"

var (
	ErrInvalidLogin = errors.New("Invalid login. Remember, everything is case-sensitive.")
	ErrUserAlreadyExists = errors.New("User with that username already exists.")
	ErrNoSuchUser = errors.New("Invalid access token")
	ErrTokenExpired = errors.New("Token has expired. Please re-log in.")
	ErrNoSuchDevice = errors.New("No device with that UUID exists.")
	ErrNegativeNumber = errors.New("Number cannot be negative")
)
