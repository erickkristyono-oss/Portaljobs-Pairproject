package job

import (
	"context"

	"job-service/internal/domain/entity"
	"job-service/internal/dto/request"
)

type JobUsecase interface {
	Create(ctx context.Context, userID uint, req request.CreateJobRequest) (*entity.Job, error)
	GetByID(ctx context.Context, id uint) (*entity.Job, error)
	GetAll(ctx context.Context) ([]*entity.Job, error)
	GetByCompanyID(ctx context.Context, userID uint) ([]*entity.Job, error)
	GetRequiredSkills(ctx context.Context, jobID uint) ([]*entity.JobRequiredSkill, error)
	Update(ctx context.Context, userID uint, id uint, req request.UpdateJobRequest) (*entity.Job, error)
	Delete(ctx context.Context, userID uint, id uint) error
	Publish(ctx context.Context, userID uint, id uint) error
	Close(ctx context.Context, userID uint, id uint) error
}
