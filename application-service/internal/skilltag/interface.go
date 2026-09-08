package skilltag

import (
	"context"

	"application-service/internal/dto/response"
)

type SkillTagUsecase interface {
	Match(
		ctx context.Context,
		userSkills []response.SkillResponse,
		jobSkills []response.JobRequiredSkillResponse,
	) SkillTagResult
}
