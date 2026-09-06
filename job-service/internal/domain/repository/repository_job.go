package repository

import (
	"context"

	"job-service/internal/domain/entity"
)

type JobRepository interface {
	Create(ctx context.Context, job *entity.Job) error
	FindByID(ctx context.Context, id uint) (*entity.Job, error)
	FindByCompanyID(ctx context.Context, companyID uint) ([]*entity.Job, error)
	FindAll(ctx context.Context) ([]*entity.Job, error)
	FindPublished(ctx context.Context) ([]*entity.Job, error)
	Update(ctx context.Context, job *entity.Job) error
	Delete(ctx context.Context, id uint) error
}
