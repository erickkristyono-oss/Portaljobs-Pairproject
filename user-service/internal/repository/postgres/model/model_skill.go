package model

import "time"

type SkillModel struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"index;not null"`
	NameLicense  string `gorm:"not null"`
	SkillTag     string `gorm:"type:varchar(150);not null"`
	Level        string `gorm:"type:varchar(20);not null"`
	Organization string
	Grade        string
	ExpiredDate  string
	Description  string
	CreatedAt    time.Time
}

func (SkillModel) TableName() string {
	return "skills"
}
