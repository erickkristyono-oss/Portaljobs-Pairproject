package repository

import (
	"context"

	"company-service/internal/domain/entity"
)

type CompanyRepository interface {
	Create(ctx context.Context, company *entity.Company) error
	Update(ctx context.Context, company *entity.Company) error
	FindByUserID(ctx context.Context, userID uint) (*entity.Company, error)
	FindByID(ctx context.Context, id uint) (*entity.Company, error)
}
