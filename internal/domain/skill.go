package domain

import (
	"context"
	"time"
)

type SkillLevel string

const (
	LevelBeginner     SkillLevel = "beginner"
	LevelIntermediate SkillLevel = "intermediate"
	LevelExpert       SkillLevel = "expert"
)

func (l SkillLevel) Valid() bool {
	switch l {
	case LevelBeginner, LevelIntermediate, LevelExpert:
		return true
	}
	return false
}

// Skill is a certificate/skill owned by a user (1-M).
type Skill struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	UserID       uint       `gorm:"index;not null" json:"user_id"`
	NameLicense  string     `gorm:"not null" json:"name_license"`
	Level        SkillLevel `gorm:"type:varchar(20)" json:"level"`
	Organization string     `json:"organization"`
	Grade        string     `json:"grade"`
	ExpiredDate  string     `json:"expired_date"`
	Description  string     `json:"description"`
	CreatedAt    time.Time  `json:"created_at"`
}

type SkillRepository interface {
	Create(ctx context.Context, s *Skill) error
	FindByUserID(ctx context.Context, userID uint) ([]Skill, error)
	FindByID(ctx context.Context, id uint) (*Skill, error)
	Delete(ctx context.Context, id uint) error
}
