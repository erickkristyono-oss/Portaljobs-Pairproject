package usecase

import (
	"context"

	"portaljob/internal/domain"
)

type AdminUsecase struct {
	users   domain.UserRepository
	reports domain.ReportRepository
}

func NewAdminUsecase(users domain.UserRepository, reports domain.ReportRepository) *AdminUsecase {
	return &AdminUsecase{users: users, reports: reports}
}

func (u *AdminUsecase) ListUsers(ctx context.Context) ([]domain.User, error) {
	return u.users.FindAll(ctx)
}

func (u *AdminUsecase) SuspendUser(ctx context.Context, id uint) error {
	if _, err := u.users.FindByID(ctx, id); err != nil {
		return err
	}
	return u.users.UpdateStatus(ctx, id, domain.StatusSuspended)
}

func (u *AdminUsecase) ListReports(ctx context.Context) ([]domain.Report, error) {
	return u.reports.FindAll(ctx)
}

// ResolveReport marks a report valid or invalid. When valid and the target is a
// user, that user is suspended.
func (u *AdminUsecase) ResolveReport(ctx context.Context, reportID uint, valid bool) (*domain.Report, error) {
	report, err := u.reports.FindByID(ctx, reportID)
	if err != nil {
		return nil, err
	}
	status := domain.ReportInvalid
	if valid {
		status = domain.ReportValid
		if report.TargetType == "user" {
			_ = u.users.UpdateStatus(ctx, report.TargetID, domain.StatusSuspended)
		}
	}
	if err := u.reports.UpdateStatus(ctx, reportID, status); err != nil {
		return nil, err
	}
	report.Status = status
	return report, nil
}
