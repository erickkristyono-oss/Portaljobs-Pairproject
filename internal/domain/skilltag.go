package domain

import "context"

// SkillTag is the shared vocabulary both sides pick from. Name is stored
// normalized (lowercase, trimmed) so matching is reliable.
type SkillTag struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
}

// UserSkillTag: skills a jobseeker HAS (M-N users<->skill_tags).
type UserSkillTag struct {
	UserID     uint `gorm:"primaryKey" json:"user_id"`
	SkillTagID uint `gorm:"primaryKey" json:"skill_tag_id"`
}

// JobSkillTag: skills a job REQUIRES (M-N jobs<->skill_tags).
type JobSkillTag struct {
	JobID      uint `gorm:"primaryKey" json:"job_id"`
	SkillTagID uint `gorm:"primaryKey" json:"skill_tag_id"`
}

// MatchWarnThreshold: below this %, we flag a warning (never a block).
const MatchWarnThreshold = 50

type MatchResult struct {
	JobID          uint     `json:"job_id"`
	Score          int      `json:"score"` // 0..100
	HasRequirement bool     `json:"has_requirement"`
	Matched        []string `json:"matched"`
	Missing        []string `json:"missing"`
	Warning        bool     `json:"warning"`
}

type SkillTagRepository interface {
	FindOrCreateByName(ctx context.Context, name string) (*SkillTag, error)
	List(ctx context.Context) ([]SkillTag, error)
	TagsByIDs(ctx context.Context, ids []uint) ([]SkillTag, error)
	SetUserTags(ctx context.Context, userID uint, tagIDs []uint) error
	UserTagIDs(ctx context.Context, userID uint) ([]uint, error)
	SetJobTags(ctx context.Context, jobID uint, tagIDs []uint) error
	JobTagIDs(ctx context.Context, jobID uint) ([]uint, error)
}
