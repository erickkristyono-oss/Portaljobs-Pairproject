package repository

import (
	"context"

	"user-service/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByID(ctx context.Context, id uint) (*entity.User, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
	FindAll(ctx context.Context) ([]entity.User, error)
}
