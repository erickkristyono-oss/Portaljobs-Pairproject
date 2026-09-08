package model

import "time"

type JobRequiredSkillModel struct {
	ID          uint      `gorm:"primaryKey"`
	JobID       uint      `gorm:"not null;index"`
	NameLicense string    `gorm:"type:varchar(150);not null"`
	SkillTag    string    `gorm:"type:varchar(150);not null"`
	Required    bool      `gorm:"not null;default:false"`
	CreatedAt   time.Time `gorm:"not null"`
}

func (JobRequiredSkillModel) TableName() string {
	return "job_required_skills"
}
