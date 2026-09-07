package usecase

import (
	"context"

	"portaljob/internal/domain"
	"portaljob/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	users domain.UserRepository
	jwt   *jwt.Manager
}

func NewAuthUsecase(users domain.UserRepository, jwtManager *jwt.Manager) *AuthUsecase {
	return &AuthUsecase{users: users, jwt: jwtManager}
}

func (u *AuthUsecase) Register(ctx context.Context, nama, email, password, role string) (*domain.User, error) {
	if !domain.Role(role).Valid() {
		return nil, domain.ErrInvalidRole
	}

	_, err := u.users.FindByEmail(ctx, email)
	if err == nil {
		return nil, domain.ErrEmailTaken
	}
	if err != domain.ErrNotFound {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Nama:     nama,
		Email:    email,
		Password: string(hash),
		Role:     domain.Role(role),
		Status:   domain.StatusActive,
	}
	if err := u.users.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *AuthUsecase) Login(ctx context.Context, email, password string) (string, *domain.User, error) {
	user, err := u.users.FindByEmail(ctx, email)
	if err == domain.ErrNotFound {
		return "", nil, domain.ErrInvalidCredential
	}
	if err != nil {
		return "", nil, err
	}
	// security: block suspended accounts from logging in
	if user.Status == domain.StatusSuspended {
		return "", nil, domain.ErrSuspended
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", nil, domain.ErrInvalidCredential
	}
	token, err := u.jwt.Generate(user.ID, string(user.Role))
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}
