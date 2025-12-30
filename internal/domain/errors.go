package domain

import (
	"errors"
)

var (
	ErrInternal       = errors.New("internal server error")
	ErrNotFound       = errors.New("resource not found")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInvalidInput   = errors.New("invalid input")
	ErrAlreadyExists  = errors.New("resource already exists")
	ErrNotImplemented = errors.New("not implemented")
	ErrNotSupported   = errors.New("not supported")
)
