package company

import (
	"context"
	"time"

	"company-service/internal/domain/entity"
	domainerrors "company-service/internal/domain/error"
	"company-service/internal/domain/repository"
	"company-service/internal/dto/response"
)

type companyUsecase struct {
	companyRepository repository.CompanyRepository
}

func NewCompanyUsecase(
	companyRepository repository.CompanyRepository,
) Usecase {

	return &companyUsecase{
		companyRepository: companyRepository,
	}
}

func (u *companyUsecase) Create(
	ctx context.Context,
	userID uint,
) (*response.CreateCompanyResponse, error) {

	_, err := u.companyRepository.FindByUserID(ctx, userID)

	if err == nil {
		return nil, domainerrors.ErrAlreadyExists
	}

	if err != domainerrors.ErrNotFound {
		return nil, err
	}

	company := &entity.Company{
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.companyRepository.Create(ctx, company); err != nil {
		return nil, err
	}

	return &response.CreateCompanyResponse{
		ID:        company.ID,
		UserID:    company.UserID,
		CreatedAt: company.CreatedAt,
	}, nil
}

func (u *companyUsecase) GetByUserID(
	ctx context.Context,
	userID uint,
) (*entity.Company, error) {

	return u.companyRepository.FindByUserID(
		ctx,
		userID,
	)
}

func (u *companyUsecase) GetByID(
	ctx context.Context,
	id uint,
) (*entity.Company, error) {

	return u.companyRepository.FindByID(
		ctx,
		id,
	)
}
