package profile

import (
	"context"

	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
)

type ProfileUsecase interface {
	Create(ctx context.Context, userID uint, req request.CreateProfileRequest) (*response.ProfileResponse, error)
	GetByUserID(ctx context.Context, userID uint) (*response.ProfileResponse, error)
	Update(ctx context.Context, userID uint, req request.UpdateProfileRequest) (*response.ProfileResponse, error)
}
