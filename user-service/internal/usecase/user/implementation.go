package user

import (
	"context"
	"user-service/internal/domain/constant"
	"user-service/internal/domain/entity"
	domainerrors "user-service/internal/domain/error"
	"user-service/internal/domain/repository"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) Register(ctx context.Context, req request.RegisterRequest) (*response.UserResponse, error) {

	existingUser, err := u.userRepository.FindByEmail(ctx, req.Email)

	if err == nil && existingUser != nil {
		return nil, domainerrors.ErrEmailTaken
	}

	role := constant.Role(req.Role)

	if !constant.IsValidRole(role) {
		return nil, domainerrors.ErrInvalidRole
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Nama:     req.Nama,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     role,
		Status:   constant.StatusActive,
	}

	if err := u.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return toUserResponse(user), nil
}

func (u *userUsecase) Login(ctx context.Context, req request.LoginRequest) (*entity.User, error) {

	user, err := u.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, domainerrors.ErrInvalidCredential
	}

	if user.Status == constant.StatusSuspended {
		return nil, domainerrors.ErrSuspended
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return nil, domainerrors.ErrInvalidCredential
	}

	return user, nil
}

func (u *userUsecase) GetByID(ctx context.Context, id uint) (*response.UserResponse, error) {

	user, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return toUserResponse(user), nil
}

func (u *userUsecase) GetAll(ctx context.Context) ([]response.UserResponse, error) {

	users, err := u.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]response.UserResponse, 0, len(users))

	for i := range users {
		result = append(result, *toUserResponse(&users[i]))
	}

	return result, nil
}

func (u *userUsecase) UpdateStatus(ctx context.Context, id uint, req request.UpdateUserStatusRequest) error {
	// Validasi status
	if !constant.IsValidStatus(req.Status) {
		return domainerrors.ErrInvalidStatus
	}

	// Pastikan user tersedia
	if _, err := u.userRepository.FindByID(ctx, id); err != nil {
		return err
	}

	// Update status user
	return u.userRepository.UpdateStatus(
		ctx,
		id,
		req.Status,
	)
}

func toUserResponse(user *entity.User) *response.UserResponse {

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
