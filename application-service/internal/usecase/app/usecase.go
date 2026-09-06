package application

import (
	"context"
	"errors"
	"time"

	"application-service/internal/client"
	"application-service/internal/domain/constant"
	"application-service/internal/domain/entity"
	domainerrors "application-service/internal/domain/error"
	"application-service/internal/domain/repository"
	"application-service/internal/dto/request"
)

type applicationUsecase struct {
	applicationRepository repository.ApplicationRepository
	jobClient             *client.JobClient
	companyClient         *client.CompanyClient
}

func NewApplicationUsecase(
	applicationRepository repository.ApplicationRepository,
	jobClient *client.JobClient,
	companyClient *client.CompanyClient,
) ApplicationUsecase {
	return &applicationUsecase{
		applicationRepository: applicationRepository,
		jobClient:             jobClient,
		companyClient:         companyClient,
	}
}

// Apply digunakan oleh jobseeker untuk melamar pekerjaan.
func (u *applicationUsecase) Apply(ctx context.Context, userID uint, req request.CreateApplicationRequest) (*entity.Application, error) {

	// 1. Cek job
	_, jobStatus, err := u.jobClient.GetJobByID(ctx, req.JobID)
	if err != nil {
		return nil, err
	}

	// 2. Job harus published
	if jobStatus != "published" {
		return nil, domainerrors.ErrConflict
	}

	// 3. Cek apakah user sudah pernah apply
	existing, err := u.applicationRepository.FindByUserAndJob(
		ctx,
		userID,
		req.JobID,
	)

	if err == nil && existing != nil {
		return nil, domainerrors.ErrConflict
	}

	if err != nil && !errors.Is(err, domainerrors.ErrNotFound) {
		return nil, err
	}

	// 4. Buat application
	application := &entity.Application{
		UserID:    userID,
		JobID:     req.JobID,
		Status:    constant.ApplicationApplied,
		AppliedAt: time.Now(),
	}

	// 5. Simpan ke database
	if err := u.applicationRepository.Create(
		ctx,
		application,
	); err != nil {
		return nil, err
	}

	return application, nil
}

// GetByID mengambil detail application.
func (u *applicationUsecase) GetByID(ctx context.Context, id uint) (*entity.Application, error) {

	return u.applicationRepository.FindByID(ctx, id)
}

// GetMyApplications mengambil semua lamaran milik jobseeker.
func (u *applicationUsecase) GetMyApplications(ctx context.Context, userID uint) ([]*entity.Application, error) {

	return u.applicationRepository.FindByUserID(
		ctx,
		userID,
	)
}

// GetApplicationsByJobID mengambil semua application
// pada job tertentu.
// Hanya company pemilik job yang boleh mengakses.
func (u *applicationUsecase) GetApplicationsByJobID(ctx context.Context, userID uint, jobID uint) ([]*entity.Application, error) {

	// 1. Ambil informasi job
	companyID, _, err := u.jobClient.GetJobByID(
		ctx,
		jobID,
	)
	if err != nil {
		return nil, err
	}

	// 2. Ambil company berdasarkan user yang login
	currentCompanyID, err := u.companyClient.GetCompanyByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	// 3. Pastikan company adalah pemilik job
	if currentCompanyID != companyID {
		return nil, domainerrors.ErrForbidden
	}

	// 4. Ambil applications
	return u.applicationRepository.FindByJobID(
		ctx,
		jobID,
	)
}

// UpdateStatus digunakan company untuk mengubah
// status application.
func (u *applicationUsecase) UpdateStatus(ctx context.Context, userID uint, id uint, status string) error {

	// 1. Validasi status
	if !constant.IsValidApplicationStatus(status) {
		return domainerrors.ErrConflict
	}

	// 2. Ambil application
	application, err := u.applicationRepository.FindByID(
		ctx,
		id,
	)
	if err != nil {
		return err
	}

	// 3. Ambil job
	companyID, _, err := u.jobClient.GetJobByID(
		ctx,
		application.JobID,
	)
	if err != nil {
		return err
	}

	// 4. Ambil company yang sedang login
	currentCompanyID, err := u.companyClient.GetCompanyByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return err
	}

	// 5. Pastikan company pemilik job
	if currentCompanyID != companyID {
		return domainerrors.ErrForbidden
	}

	// 6. Validasi transition status
	if !isValidStatusTransition(
		application.Status,
		status,
	) {
		return domainerrors.ErrConflict
	}

	// 7. Update status
	return u.applicationRepository.UpdateStatus(
		ctx,
		id,
		status,
	)
}

func (u *applicationUsecase) GetCompanyApplications(ctx context.Context, userID uint) ([]*entity.Application, error) {

	// Ambil company ID berdasarkan user yang sedang login
	companyID, err := u.companyClient.GetCompanyByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	// Ambil semua job yang dimiliki company
	jobIDs, err := u.jobClient.GetJobsByCompanyID(
		ctx,
		companyID,
	)
	if err != nil {
		return nil, err
	}

	// Ambil semua application berdasarkan job-job tersebut
	return u.applicationRepository.FindByJobIDs(
		ctx,
		jobIDs,
	)
}

// Delete digunakan untuk menghapus application milik user.
func (u *applicationUsecase) Delete(ctx context.Context, userID uint, id uint) error {

	application, err := u.applicationRepository.FindByID(
		ctx,
		id,
	)
	if err != nil {
		return err
	}

	// Hanya pemilik application yang boleh menghapus.
	if application.UserID != userID {
		return domainerrors.ErrForbidden
	}

	// Application yang sudah diterima/ditolak tidak boleh dihapus.
	if application.Status == constant.ApplicationAccepted ||
		application.Status == constant.ApplicationRejected {
		return domainerrors.ErrConflict
	}

	return u.applicationRepository.Delete(
		ctx,
		id,
	)
}

// Status transition:
//
// applied → reviewed
// reviewed → accepted
// reviewed → rejected
func isValidStatusTransition(current string, next string) bool {

	switch current {

	case constant.ApplicationApplied:
		return next == constant.ApplicationReviewed

	case constant.ApplicationReviewed:
		return next == constant.ApplicationAccepted ||
			next == constant.ApplicationRejected

	case constant.ApplicationAccepted,
		constant.ApplicationRejected:
		return false

	default:
		return false
	}
}
