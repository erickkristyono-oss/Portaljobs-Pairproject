package postgres

import (
	"context"
	"errors"

	"admin-service/internal/domain/entity"
	domainerrors "admin-service/internal/domain/error"
	"admin-service/internal/domain/repository"
	"admin-service/internal/mapper"
	"admin-service/internal/repository/postgres/model"

	"gorm.io/gorm"
)

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) repository.ReportRepository {
	return &reportRepository{
		db: db,
	}
}

func (r *reportRepository) Create(ctx context.Context, report *entity.Report) error {

	reportModel := mapper.ToReportModel(report)

	if err := r.db.WithContext(ctx).
		Create(reportModel).
		Error; err != nil {
		return err
	}

	report.ID = reportModel.ID
	report.CreatedAt = reportModel.CreatedAt
	report.UpdatedAt = reportModel.UpdatedAt

	return nil
}

func (r *reportRepository) FindAll(ctx context.Context) ([]entity.Report, error) {

	var reportModels []model.ReportModel

	if err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Find(&reportModels).
		Error; err != nil {
		return nil, err
	}

	reports := make([]entity.Report, 0, len(reportModels))

	for _, reportModel := range reportModels {
		report := mapper.ToReportEntity(&reportModel)

		reports = append(reports, *report)
	}

	return reports, nil
}

func (r *reportRepository) FindByID(ctx context.Context, id uint) (*entity.Report, error) {

	var reportModel model.ReportModel

	if err := r.db.WithContext(ctx).
		First(&reportModel, id).
		Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}

		return nil, err
	}

	return mapper.ToReportEntity(&reportModel), nil
}

func (r *reportRepository) UpdateStatus(ctx context.Context, id uint, status string) error {

	result := r.db.WithContext(ctx).
		Model(&model.ReportModel{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainerrors.ErrNotFound
	}

	return nil
}
