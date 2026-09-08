package skilltag

type MatchType string

const (
	MatchExact   MatchType = "exact"
	MatchTagOnly MatchType = "tag_only"
	MatchNone    MatchType = "none"
)

type SkillTagResult struct {
	MatchPercentage float64          `json:"match_percentage"`
	Eligible        bool             `json:"eligible"`
	SkillTags       []SkillTagDetail `json:"skill_tags"`
}

type SkillTagDetail struct {
	JobSkillID   uint      `json:"job_skill_id"`
	NameLicense  string    `json:"name_license"`
	SkillTag     string    `json:"skill_tag"`
	Required     bool      `json:"required"`
	MatchType    MatchType `json:"match_type"`
	Matched      bool      `json:"matched"`
	MatchedSkill string    `json:"matched_skill,omitempty"`
}
