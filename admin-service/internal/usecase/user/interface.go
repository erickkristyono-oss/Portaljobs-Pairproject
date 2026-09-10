package user

import (
	"context"

	"admin-service/internal/dto/request"
)

type UserUsecase interface {
	UpdateStatus(
		ctx context.Context,
		id uint,
		req request.UpdateUserStatusRequest,
	) error
}
