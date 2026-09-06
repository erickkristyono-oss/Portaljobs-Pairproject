package application

import (
	"context"

	"application-service/internal/domain/entity"
	"application-service/internal/dto/request"
)

type ApplicationUsecase interface {
	Apply(ctx context.Context, userID uint, req request.CreateApplicationRequest) (*entity.Application, error)
	GetByID(ctx context.Context, id uint) (*entity.Application, error)
	GetMyApplications(ctx context.Context, userID uint) ([]*entity.Application, error)
	GetApplicationsByJobID(ctx context.Context, userID uint, jobID uint) ([]*entity.Application, error)
	UpdateStatus(ctx context.Context, userID uint, id uint, status string) error
	GetCompanyApplications(ctx context.Context, userID uint) ([]*entity.Application, error)
	Delete(ctx context.Context, userID uint, id uint) error
}
