package repository

import (
	"context"

	"user-service/internal/domain/entity"
)

type ProfileRepository interface {
	Create(ctx context.Context, profile *entity.Profile) error
	Update(ctx context.Context, profile *entity.Profile) error
	FindByUserID(ctx context.Context, userID uint) (*entity.Profile, error)
}
