package postgres

import (
	"context"

	"user-service/internal/domain/entity"
	domainrepo "user-service/internal/domain/repository"
	"user-service/internal/mapper"
	"user-service/internal/repository/postgres/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domainrepo.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(
	ctx context.Context,
	user *entity.User,
) error {
	userModel := mapper.UserEntityToModel(user)

	if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return err
	}

	user.ID = userModel.ID
	user.CreatedAt = userModel.CreatedAt
	user.UpdatedAt = userModel.UpdatedAt

	return nil
}

func (r *userRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*entity.User, error) {
	var userModel model.UserModel

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&userModel).Error

	if err != nil {
		return nil, err
	}

	return mapper.UserModelToEntity(&userModel), nil
}

func (r *userRepository) FindByID(
	ctx context.Context,
	id uint,
) (*entity.User, error) {
	var userModel model.UserModel

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&userModel).Error

	if err != nil {
		return nil, err
	}

	return mapper.UserModelToEntity(&userModel), nil
}

func (r *userRepository) UpdateStatus(
	ctx context.Context,
	id uint,
	status string,
) error {
	result := r.db.WithContext(ctx).
		Model(&model.UserModel{}).
		Where("id = ?", id).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *userRepository) FindAll(
	ctx context.Context,
) ([]entity.User, error) {
	var userModels []model.UserModel

	if err := r.db.WithContext(ctx).
		Order("id ASC").
		Find(&userModels).Error; err != nil {
		return nil, err
	}

	users := make([]entity.User, 0, len(userModels))

	for i := range userModels {
		user := mapper.UserModelToEntity(&userModels[i])

		if user != nil {
			users = append(users, *user)
		}
	}

	return users, nil
}
