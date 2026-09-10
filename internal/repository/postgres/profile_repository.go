package postgres

import (
	"context"
	"errors"

	"portaljob/internal/domain"

	"gorm.io/gorm"
)

type profileRepository struct{ db *gorm.DB }

func NewProfileRepository(db *gorm.DB) domain.ProfileRepository { return &profileRepository{db: db} }

func (r *profileRepository) Create(ctx context.Context, p *domain.Profile) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *profileRepository) Update(ctx context.Context, p *domain.Profile) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *profileRepository) FindByUserID(ctx context.Context, userID uint) (*domain.Profile, error) {
	var p domain.Profile
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}
