package errors

import "errors"

var (
	ErrNotFound  = errors.New("company not found")
	ErrConflict  = errors.New("company already exists")
	ErrForbidden = errors.New("forbidden")
)
