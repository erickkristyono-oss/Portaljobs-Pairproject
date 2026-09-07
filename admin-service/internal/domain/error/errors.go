package error

import "errors"

var (
	ErrNotFound  = errors.New("record not found")
	ErrForbidden = errors.New("forbidden")
	ErrConflict  = errors.New("conflict")
)
