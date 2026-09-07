package domain

import "errors"

// Sentinel errors so the usecase layer can signal outcomes without knowing
// about HTTP; handlers map these to the right status codes.
var (
	ErrNotFound          = errors.New("record not found")
	ErrEmailTaken        = errors.New("email already registered")
	ErrInvalidCredential = errors.New("invalid email or password")
	ErrInvalidRole       = errors.New("invalid role")
	ErrInvalidLevel      = errors.New("invalid skill level")
	ErrForbidden         = errors.New("forbidden")
	ErrConflict          = errors.New("already exists / duplicate")
	ErrSuspended         = errors.New("account is suspended")
	ErrNoCompanyProfile  = errors.New("create your company profile first")
)
