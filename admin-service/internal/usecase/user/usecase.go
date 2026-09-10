package user

import (
	"context"

	"admin-service/internal/client"
	"admin-service/internal/dto/request"
)

type userUsecase struct {
	userClient client.UserClient
}

func NewUserUsecase(
	userClient client.UserClient,
) UserUsecase {
	return &userUsecase{
		userClient: userClient,
	}
}

func (u *userUsecase) UpdateStatus(
	ctx context.Context,
	id uint,
	req request.UpdateUserStatusRequest,
) error {

	return u.userClient.UpdateStatus(
		ctx,
		id,
		req.Status,
	)
}
