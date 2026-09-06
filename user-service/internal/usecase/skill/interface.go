package skill

import (
	"context"

	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
)

type SkillUsecase interface {
	Create(ctx context.Context, userID uint, req request.CreateSkillRequest) (*response.SkillResponse, error)
	GetByUserID(ctx context.Context, userID uint) ([]response.SkillResponse, error)
	Delete(ctx context.Context, userID uint, skillID uint) error
}
