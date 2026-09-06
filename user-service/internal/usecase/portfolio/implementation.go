package portfolio

import (
	"context"

	"user-service/internal/domain/entity"
	domainerrors "user-service/internal/domain/error"
	"user-service/internal/domain/repository"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
)

type portfolioUsecase struct {
	portfolioRepository repository.PortfolioRepository
}

func NewPortfolioUsecase(portfolioRepository repository.PortfolioRepository) PortfolioUsecase {
	return &portfolioUsecase{
		portfolioRepository: portfolioRepository,
	}
}

func (u *portfolioUsecase) Create(ctx context.Context, userID uint, req request.CreatePortfolioRequest) (*response.PortfolioResponse, error) {

	portfolio := &entity.Portfolio{
		UserID:       userID,
		NameProject:  req.NameProject,
		Organization: req.Organization,
		Description:  req.Description,
	}

	if err := u.portfolioRepository.Create(ctx, portfolio); err != nil {
		return nil, err
	}

	return toPortfolioResponse(portfolio), nil
}

func (u *portfolioUsecase) GetByUserID(ctx context.Context, userID uint) ([]response.PortfolioResponse, error) {

	portfolios, err := u.portfolioRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]response.PortfolioResponse, 0, len(portfolios))

	for i := range portfolios {
		result = append(result, *toPortfolioResponse(&portfolios[i]))
	}

	return result, nil
}

func (u *portfolioUsecase) GetByID(ctx context.Context, userID uint, portfolioID uint) (*response.PortfolioResponse, error) {

	portfolio, err := u.portfolioRepository.FindByID(ctx, portfolioID)
	if err != nil {
		return nil, err
	}

	if portfolio.UserID != userID {
		return nil, domainerrors.ErrForbidden
	}

	return toPortfolioResponse(portfolio), nil
}

func (u *portfolioUsecase) Update(ctx context.Context, userID uint, portfolioID uint, req request.UpdatePortfolioRequest) (*response.PortfolioResponse, error) {

	portfolio, err := u.portfolioRepository.FindByID(ctx, portfolioID)
	if err != nil {
		return nil, err
	}

	if portfolio.UserID != userID {
		return nil, domainerrors.ErrForbidden
	}

	portfolio.NameProject = req.NameProject
	portfolio.Organization = req.Organization
	portfolio.Description = req.Description

	if err := u.portfolioRepository.Update(ctx, portfolio); err != nil {
		return nil, err
	}

	return toPortfolioResponse(portfolio), nil
}

func (u *portfolioUsecase) Delete(ctx context.Context, userID uint, portfolioID uint) error {

	portfolio, err := u.portfolioRepository.FindByID(ctx, portfolioID)
	if err != nil {
		return err
	}

	if portfolio.UserID != userID {
		return domainerrors.ErrForbidden
	}

	return u.portfolioRepository.Delete(ctx, portfolioID)
}

func toPortfolioResponse(portfolio *entity.Portfolio) *response.PortfolioResponse {
	return &response.PortfolioResponse{
		ID:           portfolio.ID,
		UserID:       portfolio.UserID,
		NameProject:  portfolio.NameProject,
		Organization: portfolio.Organization,
		Description:  portfolio.Description,
		CreatedAt:    portfolio.CreatedAt,
		UpdatedAt:    portfolio.UpdatedAt,
	}
}
