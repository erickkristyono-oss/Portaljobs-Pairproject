package entity

import "time"

type JobRequiredSkill struct {
	ID          uint
	JobID       uint
	NameLicense string
	SkillTag    string
	Required    bool
	CreatedAt   time.Time
}
