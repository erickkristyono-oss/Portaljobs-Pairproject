package portfolio

import (
	"context"

	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
)

type PortfolioUsecase interface {
	Create(ctx context.Context, userID uint, req request.CreatePortfolioRequest) (*response.PortfolioResponse, error)
	GetByUserID(ctx context.Context, userID uint) ([]response.PortfolioResponse, error)
	GetByID(ctx context.Context, userID uint, portfolioID uint) (*response.PortfolioResponse, error)
	Update(ctx context.Context, userID uint, portfolioID uint, req request.UpdatePortfolioRequest) (*response.PortfolioResponse, error)
	Delete(ctx context.Context, userID uint, portfolioID uint) error
}
