package repository

import (
	"context"

	"user-service/internal/domain/entity"
)

type SkillRepository interface {
	Create(ctx context.Context, skill *entity.Skill) error
	FindByUserID(ctx context.Context, userID uint) ([]entity.Skill, error)
	FindByID(ctx context.Context, id uint) (*entity.Skill, error)
	Update(ctx context.Context, skill *entity.Skill) error
	Delete(ctx context.Context, id uint) error
}
