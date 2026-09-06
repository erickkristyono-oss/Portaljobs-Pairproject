package entity

import "time"

type Company struct {
	ID          uint
	UserID      uint
	Name        string
	FieldOf     string
	Address     string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
