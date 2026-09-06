package model

import "time"

type ApplicationModel struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;uniqueIndex:idx_user_job"`
	JobID     uint      `gorm:"not null;uniqueIndex:idx_user_job"`
	Status    string    `gorm:"type:varchar(20);default:'submitted'"`
	AppliedAt time.Time `gorm:"not null"`
}

func (ApplicationModel) TableName() string {
	return "applications"
}
