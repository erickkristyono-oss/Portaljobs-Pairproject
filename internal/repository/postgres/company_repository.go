package postgres

import (
	"context"
	"errors"

	"portaljob/internal/domain"

	"gorm.io/gorm"
)

type companyRepository struct{ db *gorm.DB }

func NewCompanyRepository(db *gorm.DB) domain.CompanyRepository { return &companyRepository{db: db} }

func (r *companyRepository) Create(ctx context.Context, c *domain.Company) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *companyRepository) Update(ctx context.Context, c *domain.Company) error {
	return r.db.WithContext(ctx).Save(c).Error
}

func (r *companyRepository) FindByUserID(ctx context.Context, userID uint) (*domain.Company, error) {
	var c domain.Company
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&c).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *companyRepository) FindByID(ctx context.Context, id uint) (*domain.Company, error) {
	var c domain.Company
	err := r.db.WithContext(ctx).First(&c, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}
