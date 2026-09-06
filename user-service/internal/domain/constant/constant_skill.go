package constant

type SkillLevel string

const (
	LevelBeginner     SkillLevel = "beginner"
	LevelIntermediate SkillLevel = "intermediate"
	LevelExpert       SkillLevel = "expert"
)

func IsValidSkillLevel(level SkillLevel) bool {
	switch level {
	case LevelBeginner, LevelIntermediate, LevelExpert:
		return true
	default:
		return false
	}
}
