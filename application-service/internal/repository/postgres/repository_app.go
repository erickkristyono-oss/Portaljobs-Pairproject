package postgres

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"application-service/internal/domain/entity"
	domainerrors "application-service/internal/domain/error"
	"application-service/internal/domain/repository"
	"application-service/internal/mapper"
	"application-service/internal/repository/postgres/model"
)

type applicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) repository.ApplicationRepository {
	return &applicationRepository{
		db: db,
	}
}

func (r *applicationRepository) Create(ctx context.Context, application *entity.Application) error {

	applicationModel := mapper.ToApplicationModel(application)

	if err := r.db.
		WithContext(ctx).
		Create(applicationModel).
		Error; err != nil {
		return err
	}

	application.ID = applicationModel.ID

	return nil
}

func (r *applicationRepository) FindByID(ctx context.Context, id uint) (*entity.Application, error) {

	var applicationModel model.ApplicationModel

	err := r.db.
		WithContext(ctx).
		First(&applicationModel, id).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}

		return nil, err
	}

	return mapper.ToApplicationEntity(&applicationModel), nil
}

func (r *applicationRepository) FindByUserID(ctx context.Context, userID uint) ([]*entity.Application, error) {

	var applicationModels []model.ApplicationModel

	err := r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		Order("applied_at DESC").
		Find(&applicationModels).
		Error

	if err != nil {
		return nil, err
	}

	applications := make(
		[]*entity.Application,
		0,
		len(applicationModels),
	)

	for i := range applicationModels {
		applications = append(
			applications,
			mapper.ToApplicationEntity(&applicationModels[i]),
		)
	}

	return applications, nil
}

func (r *applicationRepository) FindByJobID(ctx context.Context, jobID uint) ([]*entity.Application, error) {

	var applicationModels []model.ApplicationModel

	err := r.db.
		WithContext(ctx).
		Where("job_id = ?", jobID).
		Order("applied_at DESC").
		Find(&applicationModels).
		Error

	if err != nil {
		return nil, err
	}

	applications := make(
		[]*entity.Application,
		0,
		len(applicationModels),
	)

	for i := range applicationModels {
		applications = append(
			applications,
			mapper.ToApplicationEntity(&applicationModels[i]),
		)
	}

	return applications, nil
}

func (r *applicationRepository) FindByUserAndJob(ctx context.Context, userID uint, jobID uint) (*entity.Application, error) {

	var applicationModel model.ApplicationModel

	err := r.db.
		WithContext(ctx).
		Where(
			"user_id = ? AND job_id = ?",
			userID,
			jobID,
		).
		First(&applicationModel).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}

		return nil, err
	}

	return mapper.ToApplicationEntity(&applicationModel), nil
}

func (r *applicationRepository) UpdateStatus(ctx context.Context, id uint, status string) error {

	result := r.db.
		WithContext(ctx).
		Model(&model.ApplicationModel{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": status,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainerrors.ErrNotFound
	}

	return nil
}

func (r *applicationRepository) FindByJobIDs(ctx context.Context, jobIDs []uint) ([]*entity.Application, error) {

	if len(jobIDs) == 0 {
		return []*entity.Application{}, nil
	}

	var applicationModels []model.ApplicationModel

	err := r.db.WithContext(ctx).
		Where("job_id IN ?", jobIDs).
		Order("applied_at DESC").
		Find(&applicationModels).Error

	if err != nil {
		return nil, err
	}

	applications := make(
		[]*entity.Application,
		0,
		len(applicationModels),
	)

	for i := range applicationModels {
		applications = append(
			applications,
			mapper.ToApplicationEntity(
				&applicationModels[i],
			),
		)
	}

	return applications, nil
}

func (r *applicationRepository) Delete(ctx context.Context, id uint) error {

	result := r.db.
		WithContext(ctx).
		Delete(&model.ApplicationModel{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domainerrors.ErrNotFound
	}

	return nil
}
