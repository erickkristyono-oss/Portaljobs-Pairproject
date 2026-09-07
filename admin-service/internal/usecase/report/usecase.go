package report

import (
	"context"
	"time"

	"admin-service/internal/domain/constant"
	"admin-service/internal/domain/entity"
	domainerrors "admin-service/internal/domain/error"
	"admin-service/internal/domain/repository"
	"admin-service/internal/dto/request"
)

type reportUsecase struct {
	reportRepository repository.ReportRepository
}

func NewReportUsecase(reportRepository repository.ReportRepository) ReportUsecase {
	return &reportUsecase{
		reportRepository: reportRepository,
	}
}

func (u *reportUsecase) Create(ctx context.Context, reporterID uint, req request.CreateReportRequest) (*entity.Report, error) {

	targetType := constant.ReportTargetType(req.TargetType)

	if !constant.IsValidReportTargetType(targetType) {
		return nil, domainerrors.ErrConflict
	}

	report := &entity.Report{
		ReporterID: reporterID,
		TargetType: targetType,
		TargetID:   req.TargetID,
		Reason:     req.Reason,
		Status:     constant.ReportOpen,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := u.reportRepository.Create(ctx, report); err != nil {
		return nil, err
	}

	return report, nil
}

func (u *reportUsecase) FindAll(ctx context.Context) ([]entity.Report, error) {

	return u.reportRepository.FindAll(ctx)
}

func (u *reportUsecase) FindByID(ctx context.Context, id uint) (*entity.Report, error) {

	return u.reportRepository.FindByID(ctx, id)
}

func (u *reportUsecase) UpdateStatus(ctx context.Context, id uint, status string) error {

	report, err := u.reportRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	newStatus := constant.ReportStatus(status)

	if !constant.IsValidReportStatus(newStatus) {
		return domainerrors.ErrConflict
	}

	if !isValidTransition(report.Status, newStatus) {
		return domainerrors.ErrConflict
	}

	return u.reportRepository.UpdateStatus(
		ctx,
		id,
		status,
	)
}

func isValidTransition(current constant.ReportStatus, next constant.ReportStatus) bool {

	switch current {

	case constant.ReportOpen:
		return next == constant.ReportValid ||
			next == constant.ReportInvalid

	case constant.ReportValid:
		return next == constant.ReportClosed

	case constant.ReportInvalid:
		return next == constant.ReportClosed

	case constant.ReportClosed:
		return false

	default:
		return false
	}
}
