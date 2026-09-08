package entity

import (
	"time"
	"user-service/internal/domain/constant"
)

type Skill struct {
	ID           uint
	UserID       uint
	NameLicense  string
	Category     string
	Subcategory  string
	Level        constant.SkillLevel
	Organization string
	Grade        string
	ExpiredDate  string
	Description  string
	CreatedAt    time.Time
}
