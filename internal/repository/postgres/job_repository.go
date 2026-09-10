package postgres

import (
	"context"
	"errors"

	"portaljob/internal/domain"

	"gorm.io/gorm"
)

type jobRepository struct{ db *gorm.DB }

func NewJobRepository(db *gorm.DB) domain.JobRepository { return &jobRepository{db: db} }

func (r *jobRepository) Create(ctx context.Context, job *domain.Job) error {
	return r.db.WithContext(ctx).Create(job).Error
}

func (r *jobRepository) FindAll(ctx context.Context, f domain.JobFilter) ([]domain.Job, error) {
	var jobs []domain.Job
	q := r.db.WithContext(ctx).Model(&domain.Job{}).Where("status = ?", "published")
	if f.Lokasi != "" {
		q = q.Where("lokasi ILIKE ?", "%"+f.Lokasi+"%")
	}
	if f.Limit > 0 {
		q = q.Limit(f.Limit)
	}
	if f.Offset > 0 {
		q = q.Offset(f.Offset)
	}
	err := q.Order("created_at DESC").Find(&jobs).Error
	return jobs, err
}

func (r *jobRepository) FindByID(ctx context.Context, id uint) (*domain.Job, error) {
	var job domain.Job
	err := r.db.WithContext(ctx).First(&job, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *jobRepository) FindByCompanyID(ctx context.Context, companyID uint) ([]domain.Job, error) {
	var jobs []domain.Job
	err := r.db.WithContext(ctx).Where("company_id = ?", companyID).Order("created_at DESC").Find(&jobs).Error
	return jobs, err
}
