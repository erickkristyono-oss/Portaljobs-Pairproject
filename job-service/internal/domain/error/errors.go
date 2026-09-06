package errors

import "errors"

var (
	ErrNotFound  = errors.New("record not found")
	ErrForbidden = errors.New("forbidden")
	ErrConflict  = errors.New("already exists / duplicate")
)
