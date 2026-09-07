package domain

import "context"

// Notifier is the contract for sending notifications (e.g. WhatsApp).
// Because it is an interface, usecases can be unit-tested with a mock that
// records calls instead of hitting a real WA gateway.
type Notifier interface {
	Send(ctx context.Context, to, message string) error
}
