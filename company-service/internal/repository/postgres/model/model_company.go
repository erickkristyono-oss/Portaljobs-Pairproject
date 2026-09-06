package model

import "time"

type CompanyModel struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"uniqueIndex;not null"`
	Name        string `gorm:"not null"`
	FieldOf     string `gorm:"not null"`
	Address     string `gorm:"not null"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (CompanyModel) TableName() string {
	return "companies"
}
