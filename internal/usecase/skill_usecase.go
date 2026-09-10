package usecase

import (
	"context"

	"portaljob/internal/domain"
)

type SkillUsecase struct{ skills domain.SkillRepository }

func NewSkillUsecase(skills domain.SkillRepository) *SkillUsecase {
	return &SkillUsecase{skills: skills}
}

func (u *SkillUsecase) Add(ctx context.Context, userID uint, s domain.Skill) (*domain.Skill, error) {
	if s.Level != "" && !s.Level.Valid() {
		return nil, domain.ErrInvalidLevel
	}
	s.UserID = userID
	if err := u.skills.Create(ctx, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func (u *SkillUsecase) List(ctx context.Context, userID uint) ([]domain.Skill, error) {
	return u.skills.FindByUserID(ctx, userID)
}

// Remove deletes only if the skill belongs to userID (prevents deleting others').
func (u *SkillUsecase) Remove(ctx context.Context, userID, skillID uint) error {
	s, err := u.skills.FindByID(ctx, skillID)
	if err != nil {
		return err
	}
	if s.UserID != userID {
		return domain.ErrForbidden
	}
	return u.skills.Delete(ctx, skillID)
}
