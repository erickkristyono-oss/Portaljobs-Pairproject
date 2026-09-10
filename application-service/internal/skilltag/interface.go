package skilltag

import "application-service/internal/dto/response"

type SkillTagUsecase interface {
	Match(userSkills []response.SkillResponse, jobSkills []response.JobRequiredSkillResponse) SkillTagResult
}
