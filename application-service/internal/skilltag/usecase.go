package skilltag

import (
	"strings"

	"application-service/internal/dto/response"
)

type skillTagUsecase struct{}

func NewSkillTagUsecase() SkillTagUsecase {
	return &skillTagUsecase{}
}

func (s *skillTagUsecase) Match(
	userSkills []response.SkillResponse,
	jobSkills []response.JobRequiredSkillResponse,
) SkillTagResult {

	result := SkillTagResult{
		SkillTags: make([]SkillTagDetail, 0),
	}

	// Kalau job tidak mempunyai required skill,
	// kandidat otomatis dianggap eligible.
	if len(jobSkills) == 0 {
		result.MatchPercentage = 100
		result.Eligible = true
		return result
	}

	matchedCount := 0
	requiredMatched := true

	for _, jobSkill := range jobSkills {

		detail := SkillTagDetail{
			JobSkillID:  jobSkill.ID,
			NameLicense: jobSkill.NameLicense,
			SkillTag:    jobSkill.SkillTag,
			Required:    jobSkill.Required,
			MatchType:   MatchNone,
			Matched:     false,
		}

		// 1. EXACT MATCH
		//
		// NameLicense dan SkillTag harus sama.
		for _, userSkill := range userSkills {

			nameMatch := strings.EqualFold(
				strings.TrimSpace(userSkill.NameLicense),
				strings.TrimSpace(jobSkill.NameLicense),
			)

			tagMatch := strings.EqualFold(
				strings.TrimSpace(userSkill.SkillTag),
				strings.TrimSpace(jobSkill.SkillTag),
			)

			if nameMatch && tagMatch {
				detail.MatchType = MatchExact
				detail.Matched = true
				detail.MatchedSkill = userSkill.SkillTag
				break
			}
		}

		// 2. TAG ONLY MATCH
		//
		// Jika NameLicense tidak sama,
		// tetapi SkillTag sama.
		if !detail.Matched {

			for _, userSkill := range userSkills {

				tagMatch := strings.EqualFold(
					strings.TrimSpace(userSkill.SkillTag),
					strings.TrimSpace(jobSkill.SkillTag),
				)

				if tagMatch {
					detail.MatchType = MatchTagOnly
					detail.Matched = true
					detail.MatchedSkill = userSkill.SkillTag
					break
				}
			}
		}

		// 3. HITUNG MATCH
		if detail.Matched {
			matchedCount++
		}

		// Jika skill ini wajib tetapi
		// tidak ditemukan pada user,
		// maka user tidak eligible.
		if jobSkill.Required && !detail.Matched {
			requiredMatched = false
		}

		result.SkillTags = append(
			result.SkillTags,
			detail,
		)
	}

	// MATCH PERCENTAGE
	result.MatchPercentage =
		float64(matchedCount) /
			float64(len(jobSkills)) *
			100

	result.Eligible = requiredMatched

	return result
}
