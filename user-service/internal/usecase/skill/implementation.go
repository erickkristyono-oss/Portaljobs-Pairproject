package skill

import (
	"context"

	"user-service/internal/domain/constant"
	"user-service/internal/domain/entity"
	domainerrors "user-service/internal/domain/error"
	"user-service/internal/domain/repository"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
)

type skillUsecase struct {
	skillRepository repository.SkillRepository
}

func NewSkillUsecase(skillRepository repository.SkillRepository) SkillUsecase {
	return &skillUsecase{
		skillRepository: skillRepository,
	}
}

func (u *skillUsecase) Create(ctx context.Context, userID uint, req request.CreateSkillRequest) (*response.SkillResponse, error) {

	level := constant.SkillLevel(req.Level)

	if !constant.IsValidSkillLevel(level) {
		return nil, domainerrors.ErrInvalidLevel
	}

	skill := &entity.Skill{
		UserID:       userID,
		NameLicense:  req.NameLicense,
		SkillTag:     req.SkillTag,
		Level:        level,
		Organization: req.Organization,
		Grade:        req.Grade,
		ExpiredDate:  req.ExpiredDate,
		Description:  req.Description,
	}

	if err := u.skillRepository.Create(ctx, skill); err != nil {
		return nil, err
	}

	return toSkillResponse(skill), nil
}

func (u *skillUsecase) GetByUserID(ctx context.Context, userID uint) ([]response.SkillResponse, error) {

	skills, err := u.skillRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]response.SkillResponse, 0, len(skills))

	for i := range skills {
		result = append(result, *toSkillResponse(&skills[i]))
	}

	return result, nil
}

func (u *skillUsecase) Delete(ctx context.Context, userID uint, skillID uint) error {

	skill, err := u.skillRepository.FindByID(ctx, skillID)
	if err != nil {
		return err
	}

	if skill.UserID != userID {
		return domainerrors.ErrForbidden
	}

	return u.skillRepository.Delete(ctx, skillID)
}

func toSkillResponse(skill *entity.Skill) *response.SkillResponse {
	return &response.SkillResponse{
		ID:           skill.ID,
		UserID:       skill.UserID,
		NameLicense:  skill.NameLicense,
		SkillTag:     skill.SkillTag,
		Level:        string(skill.Level),
		Organization: skill.Organization,
		Grade:        skill.Grade,
		ExpiredDate:  skill.ExpiredDate,
		Description:  skill.Description,
		CreatedAt:    skill.CreatedAt,
	}
}
