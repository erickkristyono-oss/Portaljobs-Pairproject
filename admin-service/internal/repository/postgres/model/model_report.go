package model

import "time"

type ReportModel struct {
	ID         uint      `gorm:"primaryKey"`
	ReporterID uint      `gorm:"not null;index"`
	TargetType string    `gorm:"type:varchar(20);not null;index"`
	TargetID   uint      `gorm:"not null;index"`
	Reason     string    `gorm:"type:text;not null"`
	Status     string    `gorm:"type:varchar(20);not null;index"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
}

func (ReportModel) TableName() string {
	return "reports"
}
