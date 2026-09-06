package entity

import "time"

type Application struct {
	ID        uint
	UserID    uint
	JobID     uint
	Status    string
	AppliedAt time.Time
}
