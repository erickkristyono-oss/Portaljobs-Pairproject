package mapper

import (
	"user-service/internal/domain/constant"
	"user-service/internal/domain/entity"
	"user-service/internal/dto/response"
	"user-service/internal/repository/postgres/model"
)

func UserModelToEntity(user *model.UserModel) *entity.User {
	if user == nil {
		return nil
	}

	return &entity.User{
		ID:        user.ID,
		Nama:      user.Nama,
		Email:     user.Email,
		Password:  user.Password,
		Role:      constant.Role(user.Role),
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserEntityToModel(user *entity.User) *model.UserModel {
	if user == nil {
		return nil
	}

	return &model.UserModel{
		ID:        user.ID,
		Nama:      user.Nama,
		Email:     user.Email,
		Password:  user.Password,
		Role:      string(user.Role),
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserEntityToResponse(user *entity.User) *response.UserResponse {
	if user == nil {
		return nil
	}

	return &response.UserResponse{
		ID:        user.ID,
		Nama:      user.Nama,
		Email:     user.Email,
		Role:      string(user.Role),
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
