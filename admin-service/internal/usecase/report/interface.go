package report

import (
	"context"

	"admin-service/internal/domain/entity"
	"admin-service/internal/dto/request"
)

type ReportUsecase interface {
	Create(ctx context.Context, reporterID uint, req request.CreateReportRequest) (*entity.Report, error)
	FindAll(ctx context.Context) ([]entity.Report, error)
	FindByID(ctx context.Context, id uint) (*entity.Report, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
}
