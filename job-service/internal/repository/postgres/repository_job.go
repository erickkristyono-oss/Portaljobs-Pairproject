package postgres

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"job-service/internal/domain/entity"
	domainerrors "job-service/internal/domain/error"
	"job-service/internal/domain/repository"
	"job-service/internal/mapper"
	"job-service/internal/repository/postgres/model"
)

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) repository.JobRepository {
	return &jobRepository{
		db: db,
	}
}

func (r *jobRepository) Create(ctx context.Context, job *entity.Job) error {
	jobModel := mapper.ToJobModel(job)

	if err := r.db.WithContext(ctx).Create(jobModel).Error; err != nil {
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
		First(&jobModel, id).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}

		return nil, err
	}

	return mapper.ToJobEntity(&jobModel), nil
}

func (r *jobRepository) FindByCompanyID(ctx context.Context, companyID uint) ([]*entity.Job, error) {

	var jobModels []model.JobModel

	err := r.db.WithContext(ctx).
		Where("company_id = ?", companyID).
		Order("created_at DESC").
		Find(&jobModels).
		Error

	if err != nil {
		return nil, err
	}

	jobs := make([]*entity.Job, 0, len(jobModels))

	for i := range jobModels {
		jobs = append(jobs, mapper.ToJobEntity(&jobModels[i]))
	}

	return jobs, nil
}

func (r *jobRepository) FindAll(ctx context.Context) ([]*entity.Job, error) {

	var jobModels []model.JobModel

	err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Find(&jobModels).
		Error

	if err != nil {
		return nil, err
	}

	jobs := make([]*entity.Job, 0, len(jobModels))

	for i := range jobModels {
		jobs = append(jobs, mapper.ToJobEntity(&jobModels[i]))
	}

	return jobs, nil
}

func (r *jobRepository) Update(ctx context.Context, job *entity.Job) error {

	jobModel := mapper.ToJobModel(job)

	result := r.db.WithContext(ctx).
		Model(&model.JobModel{}).
		Where("id = ?", job.ID).
		Updates(map[string]interface{}{
			"judul":            jobModel.Judul,
			"about_role":       jobModel.AboutRole,
			"responsibilities": jobModel.Responsibilities,
			"deskripsi":        jobModel.Deskripsi,
			"lokasi":           jobModel.Lokasi,
			"gaji":             jobModel.Gaji,
			"status":           jobModel.Status,
			"updated_at":       jobModel.UpdatedAt,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainerrors.ErrNotFound
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
		return domainerrors.ErrNotFound
	}

	return nil
}
