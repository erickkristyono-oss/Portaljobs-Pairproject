package model

import "time"

type PortfolioModel struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"index;not null"`
	NameProject  string `gorm:"not null"`
	Organization string `gorm:"not null"`
	Description  string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (PortfolioModel) TableName() string {
	return "portfolios"
}
