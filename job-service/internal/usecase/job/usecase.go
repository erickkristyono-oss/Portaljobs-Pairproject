package job

import (
	"context"

	"job-service/internal/domain/constant"
	"job-service/internal/domain/entity"
	domainerror "job-service/internal/domain/error"
	"job-service/internal/domain/repository"
	"job-service/internal/dto/request"
)

type jobUsecase struct {
	jobRepository           repository.JobRepository
	requiredSkillRepository repository.JobRequiredSkillRepository
}

var _ JobUsecase = (*jobUsecase)(nil)

func NewJobUsecase(
	jobRepository repository.JobRepository,
	requiredSkillRepository repository.JobRequiredSkillRepository,
) JobUsecase {
	return &jobUsecase{
		jobRepository:           jobRepository,
		requiredSkillRepository: requiredSkillRepository,
	}
}

func (u *jobUsecase) Create(ctx context.Context, userID uint, req request.CreateJobRequest) (*entity.Job, error) {
	job := &entity.Job{
		CompanyID:        userID,
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

	requiredSkills := make(
		[]entity.JobRequiredSkill,
		0,
		len(req.RequiredSkills),
	)

	for _, skillReq := range req.RequiredSkills {
		requiredSkills = append(
			requiredSkills,
			entity.JobRequiredSkill{
				JobID:       job.ID,
				NameLicense: skillReq.NameLicense,
				SkillTag:    skillReq.SkillTag,
				Required:    skillReq.Required,
			},
		)
	}

	if err := u.requiredSkillRepository.CreateMany(
		ctx,
		requiredSkills,
	); err != nil {
		return nil, err
	}

	return job, nil
}

func (u *jobUsecase) GetByID(ctx context.Context, id uint) (*entity.Job, error) {
	return u.jobRepository.FindByID(ctx, id)
}

func (u *jobUsecase) GetAll(ctx context.Context) ([]*entity.Job, error) {
	// GET /jobs hanya menampilkan job yang published.
	return u.jobRepository.FindPublished(ctx)
}

func (u *jobUsecase) GetByCompanyID(ctx context.Context, userID uint) ([]*entity.Job, error) {
	return u.jobRepository.FindByCompanyID(ctx, userID)
}

func (u *jobUsecase) Update(ctx context.Context, userID uint, id uint, req request.UpdateJobRequest) (*entity.Job, error) {
	job, err := u.jobRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if job.CompanyID != userID {
		return nil, domainerror.ErrForbidden
	}

	job.Judul = req.Judul
	job.AboutRole = req.AboutRole
	job.Responsibilities = req.Responsibilities
	job.Deskripsi = req.Deskripsi
	job.Lokasi = req.Lokasi
	job.Gaji = req.Gaji

	if err := u.jobRepository.Update(ctx, job); err != nil {
		return nil, err
	}

	// Hapus required skill lama.
	if err := u.requiredSkillRepository.DeleteByJobID(
		ctx,
		job.ID,
	); err != nil {
		return nil, err
	}

	// Simpan required skill baru.
	requiredSkills := make(
		[]entity.JobRequiredSkill,
		0,
		len(req.RequiredSkills),
	)

	for _, skillReq := range req.RequiredSkills {
		requiredSkills = append(
			requiredSkills,
			entity.JobRequiredSkill{
				JobID:       job.ID,
				NameLicense: skillReq.NameLicense,
				SkillTag:    skillReq.SkillTag,
				Required:    skillReq.Required,
			},
		)
	}

	if err := u.requiredSkillRepository.CreateMany(
		ctx,
		requiredSkills,
	); err != nil {
		return nil, err
	}

	return job, nil
}

func (u *jobUsecase) Delete(ctx context.Context, userID uint, id uint) error {
	job, err := u.jobRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if job.CompanyID != userID {
		return domainerror.ErrForbidden
	}

	// Hapus required skills terlebih dahulu.
	if err := u.requiredSkillRepository.DeleteByJobID(
		ctx,
		id,
	); err != nil {
		return err
	}

	return u.jobRepository.Delete(ctx, id)
}

func (u *jobUsecase) Publish(ctx context.Context, userID uint, id uint) error {
	job, err := u.jobRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if job.CompanyID != userID {
		return domainerror.ErrForbidden
	}

	if job.Status != constant.JobDraft {
		return domainerror.ErrConflict
	}

	job.Status = constant.JobPublished

	return u.jobRepository.Update(ctx, job)
}

func (u *jobUsecase) Close(ctx context.Context, userID uint, id uint) error {
	job, err := u.jobRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if job.CompanyID != userID {
		return domainerror.ErrForbidden
	}

	if job.Status != constant.JobPublished {
		return domainerror.ErrConflict
	}

	job.Status = constant.JobClosed

	return u.jobRepository.Update(ctx, job)
}
