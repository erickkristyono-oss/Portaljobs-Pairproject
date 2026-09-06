package repository

import (
	"context"

	"application-service/internal/domain/entity"
)

type ApplicationRepository interface {
	Create(ctx context.Context, application *entity.Application) error
	FindByID(ctx context.Context, id uint) (*entity.Application, error)
	FindByUserID(ctx context.Context, userID uint) ([]*entity.Application, error)
	FindByJobID(ctx context.Context, jobID uint) ([]*entity.Application, error)
	FindByJobIDs(ctx context.Context, jobIDs []uint) ([]*entity.Application, error)
	FindByUserAndJob(ctx context.Context, userID uint, jobID uint) (*entity.Application, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
	Delete(ctx context.Context, id uint) error
}
