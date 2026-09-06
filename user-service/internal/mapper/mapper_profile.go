package mapper

import (
	"user-service/internal/domain/entity"
	"user-service/internal/dto/response"
	"user-service/internal/repository/postgres/model"
)

func ProfileModelToEntity(profile *model.ProfileModel) *entity.Profile {
	if profile == nil {
		return nil
	}

	return &entity.Profile{
		ID:             profile.ID,
		UserID:         profile.UserID,
		Name:           profile.Name,
		PhoneNumber:    profile.PhoneNumber,
		Address:        profile.Address,
		Faculty:        profile.Faculty,
		Major:          profile.Major,
		EducationLevel: profile.EducationLevel,
		Started:        profile.Started,
		Graduated:      profile.Graduated,
		CreatedAt:      profile.CreatedAt,
		UpdatedAt:      profile.UpdatedAt,
	}
}

func ProfileEntityToModel(profile *entity.Profile) *model.ProfileModel {
	if profile == nil {
		return nil
	}

	return &model.ProfileModel{
		ID:             profile.ID,
		UserID:         profile.UserID,
		Name:           profile.Name,
		PhoneNumber:    profile.PhoneNumber,
		Address:        profile.Address,
		Faculty:        profile.Faculty,
		Major:          profile.Major,
		EducationLevel: profile.EducationLevel,
		Started:        profile.Started,
		Graduated:      profile.Graduated,
		CreatedAt:      profile.CreatedAt,
		UpdatedAt:      profile.UpdatedAt,
	}
}

func ProfileEntityToResponse(profile *entity.Profile) *response.ProfileResponse {
	if profile == nil {
		return nil
	}

	return &response.ProfileResponse{
		ID:             profile.ID,
		UserID:         profile.UserID,
		Name:           profile.Name,
		PhoneNumber:    profile.PhoneNumber,
		Address:        profile.Address,
		Faculty:        profile.Faculty,
		Major:          profile.Major,
		EducationLevel: profile.EducationLevel,
		Started:        profile.Started,
		Graduated:      profile.Graduated,
		CreatedAt:      profile.CreatedAt,
		UpdatedAt:      profile.UpdatedAt,
	}
}
