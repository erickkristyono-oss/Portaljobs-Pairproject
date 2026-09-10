package postgres

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"company-service/internal/domain/entity"
	domainerrors "company-service/internal/domain/error"
	"company-service/internal/domain/repository"
	"company-service/internal/mapper"
	"company-service/internal/repository/postgres/model"
)

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) repository.CompanyRepository {

	return &companyRepository{
		db: db,
	}
}

func (r *companyRepository) Create(ctx context.Context, company *entity.Company) error {

	companyModel := mapper.EntityToCompanyModel(company)

	if err := r.db.WithContext(ctx).
		Create(companyModel).Error; err != nil {

		return err
	}

	*company = *mapper.CompanyModelToEntity(companyModel)

	return nil
}

func (r *companyRepository) FindByID(ctx context.Context, id uint) (*entity.Company, error) {

	var companyModel model.CompanyModel

	err := r.db.WithContext(ctx).
		First(&companyModel, id).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}

		return nil, err
	}

	return mapper.CompanyModelToEntity(&companyModel), nil
}

func (r *companyRepository) FindByUserID(ctx context.Context, userID uint) (*entity.Company, error) {

	var companyModel model.CompanyModel

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&companyModel).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}

		return nil, err
	}

	return mapper.ToCompanyEntity(&companyModel), nil
}

func (r *companyRepository) Update(ctx context.Context, company *entity.Company) error {

	result := r.db.WithContext(ctx).
		Model(&model.CompanyModel{}).
		Where("id = ?", company.ID).
		Updates(map[string]interface{}{
			"name":        company.Name,
			"phone":       company.Phone,
			"field_of":    company.FieldOf,
			"address":     company.Address,
			"description": company.Description,
			"updated_at":  company.UpdatedAt,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainerrors.ErrNotFound
	}

	return nil
}
