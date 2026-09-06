package errors

import "errors"

var (
	ErrNotFound          = errors.New("record not found")
	ErrEmailTaken        = errors.New("email already registered")
	ErrInvalidCredential = errors.New("invalid email or password")
	ErrInvalidRole       = errors.New("invalid role")
	ErrInvalidLevel      = errors.New("invalid skill level")
	ErrInvalidStatus     = errors.New("invalid status")
	ErrForbidden         = errors.New("forbidden")
	ErrConflict          = errors.New("already exists / duplicate")
	ErrSuspended         = errors.New("account is suspended")
)
