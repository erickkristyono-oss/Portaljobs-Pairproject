package repository

import (
	"context"

	"admin-service/internal/domain/entity"
)

type ReportRepository interface {
	Create(ctx context.Context, report *entity.Report) error
	FindAll(ctx context.Context) ([]entity.Report, error)
	FindByID(ctx context.Context, id uint) (*entity.Report, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
}
