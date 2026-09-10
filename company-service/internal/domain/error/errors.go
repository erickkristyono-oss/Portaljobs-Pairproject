package error

import "errors"

var (
	ErrNotFound      = errors.New("company not found")
	ErrAlreadyExists = errors.New("company already exists")
	ErrForbidden     = errors.New("forbidden")
	ErrInvalidData   = errors.New("invalid company data")
)
