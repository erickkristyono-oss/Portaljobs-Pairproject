package company_profile

import (
	"context"

	"company-service/internal/dto/request"
	"company-service/internal/dto/response"
)

type Usecase interface {
	Create(
		ctx context.Context,
		userID uint,
		req *request.CreateCompanyProfileRequest,
	) (*response.CompanyResponse, error)
	Get(ctx context.Context, userID uint) (*response.CompanyResponse, error)
	Update(
		ctx context.Context,
		userID uint,
		req *request.UpdateCompanyProfileRequest,
	) (*response.CompanyResponse, error)
}
