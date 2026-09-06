package model

import "time"

type ProfileModel struct {
	ID             uint   `gorm:"primaryKey"`
	UserID         uint   `gorm:"uniqueIndex;not null"`
	Name           string `gorm:"not null"`
	PhoneNumber    string `gorm:"not null"`
	Address        string `gorm:"not null"`
	Faculty        string
	Major          string
	EducationLevel string
	Started        int
	Graduated      int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (ProfileModel) TableName() string {
	return "profiles"
}
