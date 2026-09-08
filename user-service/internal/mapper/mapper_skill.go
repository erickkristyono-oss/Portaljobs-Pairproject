package mapper

import (
	"user-service/internal/domain/constant"
	"user-service/internal/domain/entity"
	"user-service/internal/dto/response"
	"user-service/internal/repository/postgres/model"
)

func SkillModelToEntity(skill *model.SkillModel) *entity.Skill {
	if skill == nil {
		return nil
	}

	return &entity.Skill{
		ID:           skill.ID,
		UserID:       skill.UserID,
		NameLicense:  skill.NameLicense,
		Category:     skill.Category,
		Subcategory:  skill.Subcategory,
		Level:        constant.SkillLevel(skill.Level),
		Organization: skill.Organization,
		Grade:        skill.Grade,
		ExpiredDate:  skill.ExpiredDate,
		Description:  skill.Description,
		CreatedAt:    skill.CreatedAt,
	}
}

func SkillEntityToModel(skill *entity.Skill) *model.SkillModel {
	if skill == nil {
		return nil
	}

	return &model.SkillModel{
		ID:           skill.ID,
		UserID:       skill.UserID,
		NameLicense:  skill.NameLicense,
		Category:     skill.Category,
		Subcategory:  skill.Subcategory,
		Level:        string(skill.Level),
		Organization: skill.Organization,
		Grade:        skill.Grade,
		ExpiredDate:  skill.ExpiredDate,
		Description:  skill.Description,
		CreatedAt:    skill.CreatedAt,
	}
}

func SkillEntityToResponse(skill *entity.Skill) *response.SkillResponse {
	if skill == nil {
		return nil
	}

	return &response.SkillResponse{
		ID:           skill.ID,
		UserID:       skill.UserID,
		NameLicense:  skill.NameLicense,
		Category:     skill.Category,
		Subcategory:  skill.Subcategory,
		Level:        string(skill.Level),
		Organization: skill.Organization,
		Grade:        skill.Grade,
		ExpiredDate:  skill.ExpiredDate,
		Description:  skill.Description,
		CreatedAt:    skill.CreatedAt,
	}
}
