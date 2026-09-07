package usecase

import (
	"context"

	"portaljob/internal/domain"
)

type CompanyUsecase struct{ companies domain.CompanyRepository }

func NewCompanyUsecase(companies domain.CompanyRepository) *CompanyUsecase {
	return &CompanyUsecase{companies: companies}
}

// Save creates the company profile on first call, updates it afterwards.
func (u *CompanyUsecase) Save(ctx context.Context, userID uint, name, fieldOf, address, description string) (*domain.Company, error) {
	existing, err := u.companies.FindByUserID(ctx, userID)
	if err != nil && err != domain.ErrNotFound {
		return nil, err
	}
	if err == domain.ErrNotFound {
		c := &domain.Company{UserID: userID, Name: name, FieldOf: fieldOf, Address: address, Description: description}
		if err := u.companies.Create(ctx, c); err != nil {
			return nil, err
		}
		return c, nil
	}
	existing.Name, existing.FieldOf, existing.Address, existing.Description = name, fieldOf, address, description
	if err := u.companies.Update(ctx, existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (u *CompanyUsecase) Get(ctx context.Context, userID uint) (*domain.Company, error) {
	return u.companies.FindByUserID(ctx, userID)
}
