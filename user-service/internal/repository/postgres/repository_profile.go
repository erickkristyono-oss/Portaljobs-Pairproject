package postgres

import (
	"context"
	"errors"

	"user-service/internal/domain/entity"
	domainerrors "user-service/internal/domain/error"
	domainrepo "user-service/internal/domain/repository"
	"user-service/internal/mapper"
	"user-service/internal/repository/postgres/model"

	"gorm.io/gorm"
)

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) domainrepo.ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

func (r *profileRepository) Create(
	ctx context.Context,
	profile *entity.Profile,
) error {
	profileModel := mapper.ProfileEntityToModel(profile)

	if err := r.db.WithContext(ctx).
		Create(profileModel).Error; err != nil {
		return err
	}

	profile.ID = profileModel.ID
	profile.CreatedAt = profileModel.CreatedAt
	profile.UpdatedAt = profileModel.UpdatedAt

	return nil
}

func (r *profileRepository) Update(
	ctx context.Context,
	profile *entity.Profile,
) error {
	profileModel := mapper.ProfileEntityToModel(profile)

	result := r.db.WithContext(ctx).
		Model(&model.ProfileModel{}).
		Where("id = ?", profile.ID).
		Updates(map[string]interface{}{
			"name":            profileModel.Name,
			"phone_number":    profileModel.PhoneNumber,
			"address":         profileModel.Address,
			"faculty":         profileModel.Faculty,
			"major":           profileModel.Major,
			"education_level": profileModel.EducationLevel,
			"started":         profileModel.Started,
			"graduated":       profileModel.Graduated,
			"updated_at":      profileModel.UpdatedAt,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *profileRepository) FindByUserID(
	ctx context.Context,
	userID uint,
) (*entity.Profile, error) {
	var profileModel model.ProfileModel

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&profileModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrNotFound
		}

		return nil, err
	}

	return mapper.ProfileModelToEntity(&profileModel), nil
}
