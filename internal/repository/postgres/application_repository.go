package postgres

import (
	"context"
	"errors"

	"portaljob/internal/domain"

	"gorm.io/gorm"
)

type applicationRepository struct{ db *gorm.DB }

func NewApplicationRepository(db *gorm.DB) domain.ApplicationRepository {
	return &applicationRepository{db: db}
}

func (r *applicationRepository) Create(ctx context.Context, a *domain.Application) error {
	return r.db.WithContext(ctx).Create(a).Error
}

func (r *applicationRepository) ExistsByUserAndJob(ctx context.Context, userID, jobID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&domain.Application{}).
		Where("user_id = ? AND job_id = ?", userID, jobID).Count(&count).Error
	return count > 0, err
}

func (r *applicationRepository) FindByUserID(ctx context.Context, userID uint) ([]domain.Application, error) {
	var apps []domain.Application
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("applied_at DESC").Find(&apps).Error
	return apps, err
}

func (r *applicationRepository) FindByJobID(ctx context.Context, jobID uint) ([]domain.Application, error) {
	var apps []domain.Application
	err := r.db.WithContext(ctx).Where("job_id = ?", jobID).Order("applied_at DESC").Find(&apps).Error
	return apps, err
}

func (r *applicationRepository) FindByID(ctx context.Context, id uint) (*domain.Application, error) {
	var a domain.Application
	err := r.db.WithContext(ctx).First(&a, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *applicationRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&domain.Application{}).Where("id = ?", id).Update("status", status).Error
}
