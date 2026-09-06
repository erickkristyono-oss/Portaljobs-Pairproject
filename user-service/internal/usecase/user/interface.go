package user

import (
	"context"

	"user-service/internal/domain/entity"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
)

type UserUsecase interface {
	Register(ctx context.Context, req request.RegisterRequest) (*response.UserResponse, error)
	Login(ctx context.Context, req request.LoginRequest) (*entity.User, error)
	GetByID(ctx context.Context, id uint) (*response.UserResponse, error)
	GetAll(ctx context.Context) ([]response.UserResponse, error)
	UpdateStatus(ctx context.Context, id uint, req request.UpdateUserStatusRequest) error
}
