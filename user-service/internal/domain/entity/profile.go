package entity

import "time"

type Profile struct {
	ID             uint
	UserID         uint
	Name           string
	PhoneNumber    string
	Address        string
	Faculty        string
	Major          string
	EducationLevel string
	Started        int
	Graduated      int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
