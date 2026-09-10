package model

import (
	"time"

	"job-service/internal/domain/constant"
)

type JobModel struct {
	ID               uint                    `gorm:"primaryKey"`
	CompanyID        uint                    `gorm:"not null;index"`
	Judul            string                  `gorm:"type:varchar(150);not null"`
	AboutRole        string                  `gorm:"type:text;not null"`
	Responsibilities string                  `gorm:"type:text;not null"`
	Deskripsi        string                  `gorm:"type:text;not null"`
	Lokasi           string                  `gorm:"type:varchar(150);not null"`
	Gaji             int64                   `gorm:"not null"`
	Status           constant.JobStatus      `gorm:"type:varchar(20);not null"`
	RequiredSkills   []JobRequiredSkillModel `gorm:"foreignKey:JobID;references:ID"`
	CreatedAt        time.Time               `gorm:"not null"`
	UpdatedAt        time.Time               `gorm:"not null"`
}

func (JobModel) TableName() string {
	return "jobs"
}
