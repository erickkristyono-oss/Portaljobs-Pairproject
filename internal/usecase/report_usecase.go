package usecase

import (
	"context"

	"portaljob/internal/domain"
)

type ReportUsecase struct{ reports domain.ReportRepository }

func NewReportUsecase(reports domain.ReportRepository) *ReportUsecase {
	return &ReportUsecase{reports: reports}
}

func (u *ReportUsecase) Create(ctx context.Context, reporterID uint, targetType string, targetID uint, reason string) (*domain.Report, error) {
	r := &domain.Report{
		ReporterID: reporterID,
		TargetType: targetType,
		TargetID:   targetID,
		Reason:     reason,
		Status:     domain.ReportOpen,
	}
	if err := u.reports.Create(ctx, r); err != nil {
		return nil, err
	}
	return r, nil
}
