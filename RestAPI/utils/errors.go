package utils

import "errors"

var (
	// ErrRecordNotFound ...
	ErrRecordNotFound = errors.New("record not found")
	// ErrRecordAlreadyExists ...
	ErrRecordAlreadyExists = errors.New("record already exists")
	// ErrIncorrectEmailOrPassword ...
	ErrIncorrectEmailOrPassword = errors.New("Incorrect Email or Password")
	// ErrWrongRequest ...
	ErrWrongRequest = errors.New("wrong request")
)
