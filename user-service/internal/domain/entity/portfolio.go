package entity

import "time"

type Portfolio struct {
	ID           uint
	UserID       uint
	NameProject  string
	Organization string
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
