package skilltag

import (
	"context"
	"strings"

	"application-service/internal/dto/response"
)

type skillTagUsecase struct{}

var _ SkillTagUsecase = (*skillTagUsecase)(nil)

func NewSkillTagUsecase() SkillTagUsecase {
	return &skillTagUsecase{}
}

func (u *skillTagUsecase) Match(
	ctx context.Context,
	userSkills []response.SkillResponse,
	jobSkills []response.JobRequiredSkillResponse,
) SkillTagResult {

	_ = ctx

	result := SkillTagResult{
		MatchPercentage: 0,
		Eligible:        true,
		SkillTags:       make([]SkillTagDetail, 0, len(jobSkills)),
	}

	if len(jobSkills) == 0 {
		return result
	}

	totalSkills := 0
	matchedSkills := 0
	requiredSkills := 0
	matchedRequiredSkills := 0

	for _, jobSkill := range jobSkills {

		if jobSkill.Required {
			requiredSkills++
		}

		detail := SkillTagDetail{
			JobSkillID:  jobSkill.ID,
			NameLicense: jobSkill.NameLicense,
			SkillTag:    jobSkill.SkillTag,
			Required:    jobSkill.Required,
			MatchType:   MatchNone,
			Matched:     false,
		}

		for _, userSkill := range userSkills {

			userLicense := strings.TrimSpace(
				strings.ToLower(userSkill.NameLicense),
			)

			jobLicense := strings.TrimSpace(
				strings.ToLower(jobSkill.NameLicense),
			)

			userTag := strings.TrimSpace(
				strings.ToLower(userSkill.SkillTag),
			)

			jobTag := strings.TrimSpace(
				strings.ToLower(jobSkill.SkillTag),
			)

			// Exact match:
			// NameLicense dan SkillTag sama.
			if userLicense == jobLicense &&
				userTag == jobTag {

				detail.MatchType = MatchExact
				detail.Matched = true
				detail.MatchedSkill = userSkill.NameLicense

				break
			}

			// Tag match:
			// SkillTag sama walaupun NameLicense berbeda.
			if userTag == jobTag {

				detail.MatchType = MatchTagOnly
				detail.Matched = true
				detail.MatchedSkill = userSkill.NameLicense

				break
			}
		}

		if detail.Matched {
			matchedSkills++

			if detail.Required {
				matchedRequiredSkills++
			}
		}

		if detail.Required && !detail.Matched {
			result.Eligible = false
		}

		totalSkills++

		result.SkillTags = append(
			result.SkillTags,
			detail,
		)
	}

	if totalSkills > 0 {
		result.MatchPercentage =
			float64(matchedSkills) /
				float64(totalSkills) *
				100
	}

	if requiredSkills == 0 {
		result.Eligible = true
	} else {
		result.Eligible =
			matchedRequiredSkills == requiredSkills
	}

	return result
}
