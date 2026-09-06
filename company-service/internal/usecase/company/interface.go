package company

import (
	"context"

	"company-service/internal/domain/entity"
	"company-service/internal/dto/request"
	"company-service/internal/dto/response"
)

type CompanyUsecase interface {
	Create(ctx context.Context, userID uint, req request.CreateCompanyRequest) (*response.CompanyResponse, error)
	GetByID(ctx context.Context, id uint) (*response.CompanyResponse, error)
	GetByUserID(ctx context.Context, userID uint) (*response.CompanyResponse, error)
	Update(ctx context.Context, userID uint, req request.UpdateCompanyRequest) (*response.CompanyResponse, error)
	GetCompanyByUserID(ctx context.Context, userID uint) (*entity.Company, error)
}
