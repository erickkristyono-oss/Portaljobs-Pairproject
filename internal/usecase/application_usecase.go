package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"portaljob/internal/domain"
)

type ApplicationUsecase struct {
	apps      domain.ApplicationRepository
	jobs      domain.JobRepository
	companies domain.CompanyRepository
	profiles  domain.ProfileRepository
	notifier  domain.Notifier
}

func NewApplicationUsecase(
	apps domain.ApplicationRepository,
	jobs domain.JobRepository,
	companies domain.CompanyRepository,
	profiles domain.ProfileRepository,
	notifier domain.Notifier,
) *ApplicationUsecase {
	return &ApplicationUsecase{apps: apps, jobs: jobs, companies: companies, profiles: profiles, notifier: notifier}
}

// Apply lets a jobseeker apply to a published job (once).
func (u *ApplicationUsecase) Apply(ctx context.Context, userID, jobID uint) (*domain.Application, error) {
	job, err := u.jobs.FindByID(ctx, jobID)
	if err != nil {
		return nil, err
	}
	if job.Status != "published" {
		return nil, domain.ErrNotFound
	}
	exists, err := u.apps.ExistsByUserAndJob(ctx, userID, jobID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domain.ErrConflict
	}
	app := &domain.Application{UserID: userID, JobID: jobID, Status: domain.AppSubmitted, AppliedAt: time.Now()}
	if err := u.apps.Create(ctx, app); err != nil {
		return nil, err
	}
	u.notify(ctx, userID, fmt.Sprintf("Lamaran kamu untuk lowongan #%d berhasil dikirim.", jobID))
	return app, nil
}

func (u *ApplicationUsecase) ListMine(ctx context.Context, userID uint) ([]domain.Application, error) {
	return u.apps.FindByUserID(ctx, userID)
}

// ListApplicants returns applicants for a job, only if the job belongs to the
// company owned by companyUserID.
func (u *ApplicationUsecase) ListApplicants(ctx context.Context, companyUserID, jobID uint) ([]domain.Application, error) {
	if err := u.ownsJob(ctx, companyUserID, jobID); err != nil {
		return nil, err
	}
	return u.apps.FindByJobID(ctx, jobID)
}

// Decide sets an application's status (interview/accepted/declined). Only the
// company that owns the job may decide.
func (u *ApplicationUsecase) Decide(ctx context.Context, companyUserID, appID uint, status string) (*domain.Application, error) {
	if status != domain.AppInterview && status != domain.AppAccepted && status != domain.AppDeclined {
		return nil, domain.ErrInvalidRole // reuse as "invalid input" here
	}
	app, err := u.apps.FindByID(ctx, appID)
	if err != nil {
		return nil, err
	}
	if err := u.ownsJob(ctx, companyUserID, app.JobID); err != nil {
		return nil, err
	}
	if err := u.apps.UpdateStatus(ctx, appID, status); err != nil {
		return nil, err
	}
	app.Status = status
	u.notify(ctx, app.UserID, fmt.Sprintf("Status lamaran #%d kamu sekarang: %s", appID, status))
	return app, nil
}

func (u *ApplicationUsecase) ownsJob(ctx context.Context, companyUserID, jobID uint) error {
	job, err := u.jobs.FindByID(ctx, jobID)
	if err != nil {
		return err
	}
	company, err := u.companies.FindByUserID(ctx, companyUserID)
	if err == domain.ErrNotFound {
		return domain.ErrForbidden
	}
	if err != nil {
		return err
	}
	if job.CompanyID != company.ID {
		return domain.ErrForbidden
	}
	return nil
}

// notify sends a WA message to the user's profile phone (best-effort: never
// fails the main flow). Wrap in a goroutine or a queue for production.
func (u *ApplicationUsecase) notify(ctx context.Context, userID uint, message string) {
	profile, err := u.profiles.FindByUserID(ctx, userID)
	if err != nil || profile.PhoneNumber == "" {
		return
	}
	if err := u.notifier.Send(ctx, profile.PhoneNumber, message); err != nil {
		log.Printf("[notify] gagal kirim WA ke user %d: %v", userID, err)
	}
}
