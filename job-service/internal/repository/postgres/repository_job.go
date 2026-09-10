package postgres

import (
	"context"
	"errors"

	"job-service/internal/domain/entity"
	domainerror "job-service/internal/domain/error"
	domainrepo "job-service/internal/domain/repository"
	"job-service/internal/mapper"
	"job-service/internal/repository/postgres/model"

	"gorm.io/gorm"
)

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) domainrepo.JobRepository {
	return &jobRepository{
		db: db,
	}
}

func (r *jobRepository) Create(ctx context.Context, job *entity.Job) error {
	jobModel := mapper.ToJobModel(job)

	if err := r.db.WithContext(ctx).
		Create(jobModel).Error; err != nil {
		return err
	}

	job.ID = jobModel.ID
	job.CreatedAt = jobModel.CreatedAt
	job.UpdatedAt = jobModel.UpdatedAt

	return nil
}

func (r *jobRepository) FindByID(ctx context.Context, id uint) (*entity.Job, error) {
	var jobModel model.JobModel

	err := r.db.WithContext(ctx).
		Preload("RequiredSkills").
		Where("id = ?", id).
		First(&jobModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerror.ErrNotFound
		}

		return nil, err
	}

	return mapper.ToJobEntity(&jobModel), nil
}

func (r *jobRepository) FindByCompanyID(ctx context.Context, companyID uint) ([]*entity.Job, error) {
	var jobModels []model.JobModel

	err := r.db.WithContext(ctx).
		Preload("RequiredSkills").
		Where("company_id = ?", companyID).
		Order("id DESC").
		Find(&jobModels).Error

	if err != nil {
		return nil, err
	}

	jobs := make([]*entity.Job, 0, len(jobModels))

	for i := range jobModels {
		job := mapper.ToJobEntity(&jobModels[i])

		if job != nil {
			jobs = append(jobs, job)
		}
	}

	return jobs, nil
}

func (r *jobRepository) FindAll(ctx context.Context) ([]*entity.Job, error) {
	var jobModels []model.JobModel

	err := r.db.WithContext(ctx).
		Preload("RequiredSkills").
		Order("id DESC").
		Find(&jobModels).Error

	if err != nil {
		return nil, err
	}

	jobs := make([]*entity.Job, 0, len(jobModels))

	for i := range jobModels {
		job := mapper.ToJobEntity(&jobModels[i])

		if job != nil {
			jobs = append(jobs, job)
		}
	}

	return jobs, nil
}

func (r *jobRepository) FindPublished(ctx context.Context) ([]*entity.Job, error) {
	var jobModels []model.JobModel

	err := r.db.WithContext(ctx).
		Preload("RequiredSkills").
		Where("status = ?", "published").
		Order("id DESC").
		Find(&jobModels).Error

	if err != nil {
		return nil, err
	}

	jobs := make([]*entity.Job, 0, len(jobModels))

	for i := range jobModels {
		job := mapper.ToJobEntity(&jobModels[i])

		if job != nil {
			jobs = append(jobs, job)
		}
	}

	return jobs, nil
}

func (r *jobRepository) Update(ctx context.Context, job *entity.Job) error {
	jobModel := mapper.ToJobModel(job)

	result := r.db.WithContext(ctx).
		Model(&model.JobModel{}).
		Where("id = ?", job.ID).
		Updates(jobModel)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainerror.ErrNotFound
	}

	return nil
}

func (r *jobRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).
		Delete(&model.JobModel{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainerror.ErrNotFound
	}

	return nil
}
