package repository

import (
	"context"

	"user-service/internal/domain/entity"
)

type PortfolioRepository interface {
	Create(ctx context.Context, portfolio *entity.Portfolio) error
	FindByUserID(ctx context.Context, userID uint) ([]entity.Portfolio, error)
	FindByID(ctx context.Context, id uint) (*entity.Portfolio, error)
	Update(ctx context.Context, portfolio *entity.Portfolio) error
	Delete(ctx context.Context, id uint) error
}
