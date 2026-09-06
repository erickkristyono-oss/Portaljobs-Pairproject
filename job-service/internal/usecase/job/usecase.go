package job

import (
	"context"
	"time"

	"job-service/internal/domain/constant"
	"job-service/internal/domain/entity"
	domainerrors "job-service/internal/domain/error"
	"job-service/internal/domain/repository"
	"job-service/internal/dto/request"
)

type CompanyClient interface {
	GetCompanyByUserID(ctx context.Context, userID uint) (uint, error)
}

type jobUsecase struct {
	jobRepository repository.JobRepository
	companyClient CompanyClient
}

func NewJobUsecase(jobRepository repository.JobRepository, companyClient CompanyClient) JobUsecase {
	return &jobUsecase{
		jobRepository: jobRepository,
		companyClient: companyClient,
	}
}

func (u *jobUsecase) Create(ctx context.Context, userID uint, req request.CreateJobRequest) (*entity.Job, error) {

	companyID, err := u.companyClient.GetCompanyByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	job := &entity.Job{
		CompanyID:        companyID,
		Judul:            req.Judul,
		AboutRole:        req.AboutRole,
		Responsibilities: req.Responsibilities,
		Deskripsi:        req.Deskripsi,
		Lokasi:           req.Lokasi,
		Gaji:             req.Gaji,
		Status:           constant.JobDraft,
	}

	if err := u.jobRepository.Create(ctx, job); err != nil {
		return nil, err
	}

	return job, nil
}

func (u *jobUsecase) GetByID(ctx context.Context, id uint) (*entity.Job, error) {

	return u.jobRepository.FindByID(ctx, id)
}

func (u *jobUsecase) GetAll(ctx context.Context) ([]*entity.Job, error) {

	return u.jobRepository.FindPublished(ctx)
}

func (u *jobUsecase) GetByCompanyID(ctx context.Context, userID uint) ([]*entity.Job, error) {

	companyID, err := u.companyClient.GetCompanyByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return u.jobRepository.FindByCompanyID(ctx, companyID)
}

func (u *jobUsecase) Update(ctx context.Context, userID uint, id uint, req request.UpdateJobRequest) (*entity.Job, error) {

	job, err := u.jobRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := u.checkOwnership(ctx, userID, job.CompanyID); err != nil {
		return nil, err
	}

	// Job yang sudah closed tidak boleh diedit.
	if job.Status == constant.JobClosed {
		return nil, domainerrors.ErrConflict
	}

	job.Judul = req.Judul
	job.AboutRole = req.AboutRole
	job.Responsibilities = req.Responsibilities
	job.Deskripsi = req.Deskripsi
	job.Lokasi = req.Lokasi
	job.Gaji = req.Gaji
	job.UpdatedAt = time.Now()

	if err := u.jobRepository.Update(ctx, job); err != nil {
		return nil, err
	}

	return job, nil
}

func (u *jobUsecase) Delete(ctx context.Context, userID uint, id uint) error {

	job, err := u.jobRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if err := u.checkOwnership(ctx, userID, job.CompanyID); err != nil {
		return err
	}

	if job.Status != constant.JobDraft {
		return domainerrors.ErrConflict
	}

	return u.jobRepository.Delete(ctx, id)
}

func (u *jobUsecase) Publish(ctx context.Context, userID uint, id uint) error {

	job, err := u.jobRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if err := u.checkOwnership(ctx, userID, job.CompanyID); err != nil {
		return err
	}

	if job.Status != constant.JobDraft {
		return domainerrors.ErrConflict
	}

	job.Status = constant.JobPublished
	job.UpdatedAt = time.Now()

	return u.jobRepository.Update(ctx, job)
}

func (u *jobUsecase) Close(ctx context.Context, userID uint, id uint) error {

	job, err := u.jobRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if err := u.checkOwnership(ctx, userID, job.CompanyID); err != nil {
		return err
	}

	if job.Status != constant.JobPublished {
		return domainerrors.ErrConflict
	}

	job.Status = constant.JobClosed
	job.UpdatedAt = time.Now()

	return u.jobRepository.Update(ctx, job)
}

func (u *jobUsecase) checkOwnership(ctx context.Context, userID uint, companyID uint) error {

	currentCompanyID, err := u.companyClient.GetCompanyByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if currentCompanyID != companyID {
		return domainerrors.ErrForbidden
	}

	return nil
}
