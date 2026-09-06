package repository

import (
	"context"

	"job-service/internal/domain/entity"
)

type JobFilter struct {
	Lokasi string
	Limit  int
	Offset int
}

type JobRepository interface {
	Create(ctx context.Context, job *entity.Job) error
	Update(ctx context.Context, job *entity.Job) error
	Delete(ctx context.Context, id uint) error

	FindAll(ctx context.Context, filter JobFilter) ([]entity.Job, error)
	FindByID(ctx context.Context, id uint) (*entity.Job, error)
	FindByCompanyID(ctx context.Context, companyID uint) ([]entity.Job, error)
}
