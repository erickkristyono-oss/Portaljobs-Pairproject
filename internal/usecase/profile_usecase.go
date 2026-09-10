package usecase

import (
	"context"

	"portaljob/internal/domain"
)

type ProfileUsecase struct{ profiles domain.ProfileRepository }

func NewProfileUsecase(profiles domain.ProfileRepository) *ProfileUsecase {
	return &ProfileUsecase{profiles: profiles}
}

func (u *ProfileUsecase) Save(ctx context.Context, userID uint, p domain.Profile) (*domain.Profile, error) {
	existing, err := u.profiles.FindByUserID(ctx, userID)
	if err != nil && err != domain.ErrNotFound {
		return nil, err
	}
	if err == domain.ErrNotFound {
		p.UserID = userID
		if err := u.profiles.Create(ctx, &p); err != nil {
			return nil, err
		}
		return &p, nil
	}
	p.ID, p.UserID = existing.ID, userID
	if err := u.profiles.Update(ctx, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (u *ProfileUsecase) Get(ctx context.Context, userID uint) (*domain.Profile, error) {
	return u.profiles.FindByUserID(ctx, userID)
}
