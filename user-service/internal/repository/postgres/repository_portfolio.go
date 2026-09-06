package postgres

import (
	"context"

	"user-service/internal/domain/entity"
	domainrepo "user-service/internal/domain/repository"
	"user-service/internal/mapper"
	"user-service/internal/repository/postgres/model"

	"gorm.io/gorm"
)

type portfolioRepository struct {
	db *gorm.DB
}

func NewPortfolioRepository(db *gorm.DB) domainrepo.PortfolioRepository {
	return &portfolioRepository{
		db: db,
	}
}

func (r *portfolioRepository) Create(
	ctx context.Context,
	portfolio *entity.Portfolio,
) error {
	portfolioModel := mapper.PortfolioEntityToModel(portfolio)

	if err := r.db.WithContext(ctx).
		Create(portfolioModel).Error; err != nil {
		return err
	}

	portfolio.ID = portfolioModel.ID
	portfolio.CreatedAt = portfolioModel.CreatedAt
	portfolio.UpdatedAt = portfolioModel.UpdatedAt

	return nil
}

func (r *portfolioRepository) FindByUserID(
	ctx context.Context,
	userID uint,
) ([]entity.Portfolio, error) {
	var portfolioModels []model.PortfolioModel

	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("id ASC").
		Find(&portfolioModels).Error; err != nil {
		return nil, err
	}

	portfolios := make([]entity.Portfolio, 0, len(portfolioModels))

	for i := range portfolioModels {
		portfolio := mapper.PortfolioModelToEntity(
			&portfolioModels[i],
		)

		if portfolio != nil {
			portfolios = append(portfolios, *portfolio)
		}
	}

	return portfolios, nil
}

func (r *portfolioRepository) FindByID(
	ctx context.Context,
	id uint,
) (*entity.Portfolio, error) {
	var portfolioModel model.PortfolioModel

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&portfolioModel).Error

	if err != nil {
		return nil, err
	}

	return mapper.PortfolioModelToEntity(&portfolioModel), nil
}

func (r *portfolioRepository) Update(
	ctx context.Context,
	portfolio *entity.Portfolio,
) error {
	portfolioModel := mapper.PortfolioEntityToModel(portfolio)

	result := r.db.WithContext(ctx).
		Model(&model.PortfolioModel{}).
		Where("id = ?", portfolio.ID).
		Updates(map[string]interface{}{
			"name_project": portfolioModel.NameProject,
			"organization": portfolioModel.Organization,
			"description":  portfolioModel.Description,
			"updated_at":   portfolioModel.UpdatedAt,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *portfolioRepository) Delete(
	ctx context.Context,
	id uint,
) error {
	result := r.db.WithContext(ctx).
		Delete(&model.PortfolioModel{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
