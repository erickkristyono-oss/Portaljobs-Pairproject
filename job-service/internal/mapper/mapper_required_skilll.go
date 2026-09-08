package mapper

import (
	"job-service/internal/domain/entity"
	"job-service/internal/dto/response"
	"job-service/internal/repository/postgres/model"
)

func ToJobRequiredSkillEntity(skillModel *model.JobRequiredSkillModel) *entity.JobRequiredSkill {
	if skillModel == nil {
		return nil
	}

	return &entity.JobRequiredSkill{
		ID:          skillModel.ID,
		JobID:       skillModel.JobID,
		NameLicense: skillModel.NameLicense,
		SkillTag:    skillModel.SkillTag,
		Required:    skillModel.Required,
		CreatedAt:   skillModel.CreatedAt,
	}
}

func ToJobRequiredSkillModel(skill *entity.JobRequiredSkill) *model.JobRequiredSkillModel {
	if skill == nil {
		return nil
	}

	return &model.JobRequiredSkillModel{
		ID:          skill.ID,
		JobID:       skill.JobID,
		NameLicense: skill.NameLicense,
		SkillTag:    skill.SkillTag,
		Required:    skill.Required,
		CreatedAt:   skill.CreatedAt,
	}
}

func ToJobRequiredSkillResponse(skill *entity.JobRequiredSkill) response.JobRequiredSkillResponse {
	return response.JobRequiredSkillResponse{
		ID:          skill.ID,
		JobID:       skill.JobID,
		NameLicense: skill.NameLicense,
		SkillTag:    skill.SkillTag,
		Required:    skill.Required,
		CreatedAt:   skill.CreatedAt,
	}
}
