package company

import (
	"context"

	"company-service/internal/domain/entity"
	"company-service/internal/dto/response"
)

type Usecase interface {
	Create(ctx context.Context, userID uint) (*response.CreateCompanyResponse, error)
	GetByUserID(
		ctx context.Context,
		userID uint,
	) (*entity.Company, error)
	GetByID(
		ctx context.Context,
		id uint,
	) (*entity.Company, error)
}
