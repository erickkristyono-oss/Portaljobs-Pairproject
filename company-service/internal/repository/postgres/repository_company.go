package postgres

import (
	"context"
	"errors"

	"company-service/internal/domain/entity"
	domainerrors "company-service/internal/domain/error"
	domainrepo "company-service/internal/domain/repository"
	"company-service/internal/mapper"
	"company-service/internal/repository/postgres/model"

	"gorm.io/gorm"
)

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) domainrepo.CompanyRepository {
	return &companyRepository{
		db: db,
	}
}

func (r *companyRepository) Create(
	ctx context.Context,
	company *entity.Company,
) error {
	companyModel := mapper.CompanyEntityToModel(company)

	if err := r.db.WithContext(ctx).
		Create(companyModel).Error; err != nil {
		return err
	}

	company.ID = companyModel.ID
	company.CreatedAt = companyModel.CreatedAt
	company.UpdatedAt = companyModel.UpdatedAt

	return nil
}

func (r *companyRepository) FindByID(
	ctx context.Context,
	id uint,
) (*entity.Company, error) {
	var companyModel model.CompanyModel

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&companyModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}

		return nil, err
	}

	return mapper.CompanyModelToEntity(&companyModel), nil
}

func (r *companyRepository) FindByUserID(
	ctx context.Context,
	userID uint,
) (*entity.Company, error) {
	var companyModel model.CompanyModel

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&companyModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}

		return nil, err
	}

	return mapper.CompanyModelToEntity(&companyModel), nil
}

func (r *companyRepository) Update(
	ctx context.Context,
	company *entity.Company,
) error {
	companyModel := mapper.CompanyEntityToModel(company)

	result := r.db.WithContext(ctx).
		Model(&model.CompanyModel{}).
		Where("id = ?", company.ID).
		Updates(companyModel)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainerrors.ErrNotFound
	}

	return nil
}
