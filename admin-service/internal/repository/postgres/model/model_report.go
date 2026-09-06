package model

import "time"

type ReportModel struct {
	ID         uint   `gorm:"primaryKey"`
	ReporterID uint   `gorm:"index;not null"`
	TargetType string `gorm:"type:varchar(20);not null"`
	TargetID   uint   `gorm:"index;not null"`
	Reason     string `gorm:"not null"`
	Status     string `gorm:"type:varchar(20);default:'open';not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (ReportModel) TableName() string {
	return "reports"
}
