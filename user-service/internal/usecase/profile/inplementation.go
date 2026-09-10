package profile

import (
	"context"

	"user-service/internal/domain/entity"
	domainerrors "user-service/internal/domain/error"
	"user-service/internal/domain/repository"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
)

type profileUsecase struct {
	profileRepository repository.ProfileRepository
}

func NewProfileUsecase(profileRepository repository.ProfileRepository) ProfileUsecase {
	return &profileUsecase{
		profileRepository: profileRepository,
	}
}

func (u *profileUsecase) Create(ctx context.Context, userID uint, req request.CreateProfileRequest) (*response.ProfileResponse, error) {

	existing, err := u.profileRepository.FindByUserID(ctx, userID)
	if err == nil && existing != nil {
		return nil, domainerrors.ErrConflict
	}

	if err != nil && err != domainerrors.ErrNotFound {
		return nil, err
	}

	profile := &entity.Profile{
		UserID:         userID,
		Name:           req.Name,
		PhoneNumber:    req.PhoneNumber,
		Address:        req.Address,
		Faculty:        req.Faculty,
		Major:          req.Major,
		EducationLevel: req.EducationLevel,
		Started:        req.Started,
		Graduated:      req.Graduated,
	}

	if err := u.profileRepository.Create(ctx, profile); err != nil {
		return nil, err
	}

	return toProfileResponse(profile), nil
}

func (u *profileUsecase) GetByUserID(ctx context.Context, userID uint) (*response.ProfileResponse, error) {

	profile, err := u.profileRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return toProfileResponse(profile), nil
}

func (u *profileUsecase) Update(ctx context.Context, userID uint, req request.UpdateProfileRequest) (*response.ProfileResponse, error) {

	profile, err := u.profileRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	profile.Name = req.Name
	profile.PhoneNumber = req.PhoneNumber
	profile.Address = req.Address
	profile.Faculty = req.Faculty
	profile.Major = req.Major
	profile.EducationLevel = req.EducationLevel
	profile.Started = req.Started
	profile.Graduated = req.Graduated
	profile.CVURL = req.CVURL
	profile.PortfolioURL = req.PortfolioURL

	if err := u.profileRepository.Update(ctx, profile); err != nil {
		return nil, err
	}

	return toProfileResponse(profile), nil
}

func toProfileResponse(profile *entity.Profile) *response.ProfileResponse {
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
		CVURL:          profile.CVURL,
		PortfolioURL:   profile.PortfolioURL,
		CreatedAt:      profile.CreatedAt,
		UpdatedAt:      profile.UpdatedAt,
	}
}
