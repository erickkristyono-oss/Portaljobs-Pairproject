package company

import (
	"context"
	"errors"

	"company-service/internal/domain/entity"
	domainerrors "company-service/internal/domain/error"
	"company-service/internal/domain/repository"
	"company-service/internal/dto/request"
	"company-service/internal/dto/response"
	"company-service/internal/mapper"
)

type companyUsecase struct {
	companyRepository repository.CompanyRepository
}

func NewCompanyUsecase(
	companyRepository repository.CompanyRepository,
) CompanyUsecase {
	return &companyUsecase{
		companyRepository: companyRepository,
	}
}

// Create membuat company profile baru untuk user.
func (u *companyUsecase) Create(
	ctx context.Context,
	userID uint,
	req request.CreateCompanyRequest,
) (*response.CompanyResponse, error) {

	// Satu user hanya boleh memiliki satu company.
	existing, err := u.companyRepository.FindByUserID(ctx, userID)

	if err == nil && existing != nil {
		return nil, domainerrors.ErrConflict
	}

	if err != nil && !errors.Is(err, domainerrors.ErrNotFound) {
		return nil, err
	}

	company := &entity.Company{
		UserID:      userID,
		Name:        req.Name,
		FieldOf:     req.FieldOf,
		Address:     req.Address,
		Description: req.Description,
	}

	if err := u.companyRepository.Create(ctx, company); err != nil {
		return nil, err
	}

	return mapper.CompanyEntityToResponse(company), nil
}

// GetByID mengambil company berdasarkan ID.
func (u *companyUsecase) GetByID(
	ctx context.Context,
	id uint,
) (*response.CompanyResponse, error) {

	company, err := u.companyRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.CompanyEntityToResponse(company), nil
}

// GetByUserID mengambil company berdasarkan user ID.
func (u *companyUsecase) GetByUserID(
	ctx context.Context,
	userID uint,
) (*response.CompanyResponse, error) {

	company, err := u.companyRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return mapper.CompanyEntityToResponse(company), nil
}

// Update mengubah company milik user yang sedang login.
func (u *companyUsecase) Update(
	ctx context.Context,
	userID uint,
	req request.UpdateCompanyRequest,
) (*response.CompanyResponse, error) {

	company, err := u.companyRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Ownership terjamin karena company dicari berdasarkan
	// userID yang berasal dari JWT.
	company.Name = req.Name
	company.FieldOf = req.FieldOf
	company.Address = req.Address
	company.Description = req.Description

	if err := u.companyRepository.Update(ctx, company); err != nil {
		return nil, err
	}

	return mapper.CompanyEntityToResponse(company), nil
}

// GetCompanyByUserID mengambil company berdasarkan user ID
// untuk kebutuhan komunikasi internal antar service.
func (u *companyUsecase) GetCompanyByUserID(
	ctx context.Context,
	userID uint,
) (*entity.Company, error) {

	return u.companyRepository.FindByUserID(ctx, userID)
}
