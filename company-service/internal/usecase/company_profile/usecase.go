package company_profile

import (
	"context"
	"errors"
	"time"

	"company-service/internal/domain/entity"
	domainerrors "company-service/internal/domain/error"
	"company-service/internal/domain/repository"
	"company-service/internal/dto/request"
	"company-service/internal/dto/response"
)

type companyProfileUsecase struct {
	companyRepository repository.CompanyRepository
}

func NewCompanyProfileUsecase(
	companyRepository repository.CompanyRepository,
) Usecase {
	return &companyProfileUsecase{
		companyRepository: companyRepository,
	}
}

func (u *companyProfileUsecase) Create(
	ctx context.Context,
	userID uint,
	req *request.CreateCompanyProfileRequest,
) (*response.CompanyResponse, error) {

	// Cek apakah company profile sudah ada
	existingCompany, err := u.companyRepository.FindByUserID(ctx, userID)

	if err == nil && existingCompany != nil {
		return nil, domainerrors.ErrAlreadyExists
	}

	// Jika error bukan karena data tidak ditemukan,
	// kembalikan error tersebut
	if err != nil && !errors.Is(err, domainerrors.ErrNotFound) {
		return nil, err
	}

	company := &entity.Company{
		UserID:      userID,
		Name:        req.Name,
		Phone:       req.Phone,
		FieldOf:     req.FieldOf,
		Address:     req.Address,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := u.companyRepository.Create(ctx, company); err != nil {
		return nil, err
	}

	return &response.CompanyResponse{
		ID:          company.ID,
		UserID:      company.UserID,
		Name:        company.Name,
		Phone:       company.Phone,
		FieldOf:     company.FieldOf,
		Address:     company.Address,
		Description: company.Description,
		CreatedAt:   company.CreatedAt,
		UpdatedAt:   company.UpdatedAt,
	}, nil
}

func (u *companyProfileUsecase) Get(
	ctx context.Context,
	userID uint,
) (*response.CompanyResponse, error) {

	company, err := u.companyRepository.FindByUserID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return &response.CompanyResponse{
		ID:          company.ID,
		UserID:      company.UserID,
		Name:        company.Name,
		Phone:       company.Phone,
		FieldOf:     company.FieldOf,
		Address:     company.Address,
		Description: company.Description,
		CreatedAt:   company.CreatedAt,
		UpdatedAt:   company.UpdatedAt,
	}, nil
}

func (u *companyProfileUsecase) Update(
	ctx context.Context,
	userID uint,
	req *request.UpdateCompanyProfileRequest,
) (*response.CompanyResponse, error) {

	company, err := u.companyRepository.FindByUserID(ctx, userID)

	if err != nil {
		return nil, err
	}

	company.Name = req.Name
	company.Phone = req.Phone
	company.FieldOf = req.FieldOf
	company.Address = req.Address
	company.Description = req.Description
	company.UpdatedAt = time.Now()

	if err := u.companyRepository.Update(ctx, company); err != nil {
		return nil, err
	}

	return &response.CompanyResponse{
		ID:          company.ID,
		UserID:      company.UserID,
		Name:        company.Name,
		Phone:       company.Phone,
		FieldOf:     company.FieldOf,
		Address:     company.Address,
		Description: company.Description,
		CreatedAt:   company.CreatedAt,
		UpdatedAt:   company.UpdatedAt,
	}, nil
}
