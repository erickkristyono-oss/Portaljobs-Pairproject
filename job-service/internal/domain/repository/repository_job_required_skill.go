package repository

import (
	"context"

	"job-service/internal/domain/entity"
)

type JobRequiredSkillRepository interface {
	Create(ctx context.Context, skill *entity.JobRequiredSkill) error
	CreateMany(ctx context.Context, skills []entity.JobRequiredSkill) error
	FindByJobID(ctx context.Context, jobID uint) ([]entity.JobRequiredSkill, error)
	DeleteByJobID(ctx context.Context, jobID uint) error
}
