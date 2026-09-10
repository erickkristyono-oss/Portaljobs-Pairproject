package postgres

import (
	"context"
	"errors"

	"portaljob/internal/domain"

	"gorm.io/gorm"
)

type reportRepository struct{ db *gorm.DB }

func NewReportRepository(db *gorm.DB) domain.ReportRepository { return &reportRepository{db: db} }

func (r *reportRepository) Create(ctx context.Context, rep *domain.Report) error {
	return r.db.WithContext(ctx).Create(rep).Error
}

func (r *reportRepository) FindAll(ctx context.Context) ([]domain.Report, error) {
	var reports []domain.Report
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&reports).Error
	return reports, err
}

func (r *reportRepository) FindByID(ctx context.Context, id uint) (*domain.Report, error) {
	var rep domain.Report
	err := r.db.WithContext(ctx).First(&rep, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &rep, nil
}

func (r *reportRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return r.db.WithContext(ctx).Model(&domain.Report{}).Where("id = ?", id).Update("status", status).Error
}
