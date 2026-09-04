package postgres

import (
	"context"
	"errors"

	"portaljob/internal/domain"

	"gorm.io/gorm"
)

type portfolioRepository struct{ db *gorm.DB }

func NewPortfolioRepository(db *gorm.DB) domain.PortfolioRepository {
	return &portfolioRepository{db: db}
}

func (r *portfolioRepository) Create(ctx context.Context, p *domain.Portfolio) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *portfolioRepository) FindByUserID(ctx context.Context, userID uint) ([]domain.Portfolio, error) {
	var items []domain.Portfolio
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("id DESC").Find(&items).Error
	return items, err
}

func (r *portfolioRepository) FindByID(ctx context.Context, id uint) (*domain.Portfolio, error) {
	var p domain.Portfolio
	err := r.db.WithContext(ctx).First(&p, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *portfolioRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Portfolio{}, id).Error
}
