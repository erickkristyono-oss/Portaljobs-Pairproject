package model

import "time"

type ApplicationModel struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;index"`
	JobID     uint      `gorm:"not null;index"`
	Status    string    `gorm:"type:varchar(20);not null;index"`
	AppliedAt time.Time `gorm:"not null"`
}

func (ApplicationModel) TableName() string {
	return "applications"
}
