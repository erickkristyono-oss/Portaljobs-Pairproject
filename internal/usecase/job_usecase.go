package usecase

import (
	"context"

	"portaljob/internal/domain"
)

type JobUsecase struct {
	jobs      domain.JobRepository
	companies domain.CompanyRepository
	tags      domain.SkillTagRepository
}

func NewJobUsecase(jobs domain.JobRepository, companies domain.CompanyRepository, tags domain.SkillTagRepository) *JobUsecase {
	return &JobUsecase{jobs: jobs, companies: companies, tags: tags}
}

func (u *JobUsecase) Create(ctx context.Context, userID uint, judul, aboutRole, responsibilities, deskripsi, lokasi string, gaji int64, requiredSkills []string) (*domain.Job, error) {
	company, err := u.companies.FindByUserID(ctx, userID)
	if err == domain.ErrNotFound {
		return nil, domain.ErrNoCompanyProfile
	}
	if err != nil {
		return nil, err
	}
	job := &domain.Job{
		CompanyID:        company.ID,
		Judul:            judul,
		AboutRole:        aboutRole,
		Responsibilities: responsibilities,
		Deskripsi:        deskripsi,
		Lokasi:           lokasi,
		Gaji:             gaji,
		Status:           "published",
	}
	if err := u.jobs.Create(ctx, job); err != nil {
		return nil, err
	}
	if len(requiredSkills) > 0 {
		ids, err := resolveTagIDs(ctx, u.tags, requiredSkills)
		if err != nil {
			return nil, err
		}
		if err := u.tags.SetJobTags(ctx, job.ID, ids); err != nil {
			return nil, err
		}
	}
	return job, nil
}

func (u *JobUsecase) List(ctx context.Context, lokasi string, limit, offset int) ([]domain.Job, error) {
	return u.jobs.FindAll(ctx, domain.JobFilter{Lokasi: lokasi, Limit: limit, Offset: offset})
}

func (u *JobUsecase) Detail(ctx context.Context, id uint) (*domain.Job, error) {
	return u.jobs.FindByID(ctx, id)
}

func (u *JobUsecase) ListMine(ctx context.Context, userID uint) ([]domain.Job, error) {
	company, err := u.companies.FindByUserID(ctx, userID)
	if err == domain.ErrNotFound {
		return nil, domain.ErrNoCompanyProfile
	}
	if err != nil {
		return nil, err
	}
	return u.jobs.FindByCompanyID(ctx, company.ID)
}
