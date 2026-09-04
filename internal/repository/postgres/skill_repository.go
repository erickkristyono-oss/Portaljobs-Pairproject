package postgres

import (
	"context"
	"errors"

	"portaljob/internal/domain"

	"gorm.io/gorm"
)

type skillRepository struct{ db *gorm.DB }

func NewSkillRepository(db *gorm.DB) domain.SkillRepository { return &skillRepository{db: db} }

func (r *skillRepository) Create(ctx context.Context, s *domain.Skill) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *skillRepository) FindByUserID(ctx context.Context, userID uint) ([]domain.Skill, error) {
	var skills []domain.Skill
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("id DESC").Find(&skills).Error
	return skills, err
}

func (r *skillRepository) FindByID(ctx context.Context, id uint) (*domain.Skill, error) {
	var s domain.Skill
	err := r.db.WithContext(ctx).First(&s, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *skillRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Skill{}, id).Error
}
