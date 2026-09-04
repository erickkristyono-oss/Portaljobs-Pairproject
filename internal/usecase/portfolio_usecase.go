package usecase

import (
	"context"

	"portaljob/internal/domain"
)

type PortfolioUsecase struct{ portfolios domain.PortfolioRepository }

func NewPortfolioUsecase(portfolios domain.PortfolioRepository) *PortfolioUsecase {
	return &PortfolioUsecase{portfolios: portfolios}
}

func (u *PortfolioUsecase) Add(ctx context.Context, userID uint, p domain.Portfolio) (*domain.Portfolio, error) {
	p.UserID = userID
	if err := u.portfolios.Create(ctx, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (u *PortfolioUsecase) List(ctx context.Context, userID uint) ([]domain.Portfolio, error) {
	return u.portfolios.FindByUserID(ctx, userID)
}

func (u *PortfolioUsecase) Remove(ctx context.Context, userID, id uint) error {
	p, err := u.portfolios.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if p.UserID != userID {
		return domain.ErrForbidden
	}
	return u.portfolios.Delete(ctx, id)
}
