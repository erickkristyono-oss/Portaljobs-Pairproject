package domain

import (
	"context"
	"time"
)

// Profile is the jobseeker profile, linked 1-1 to a user.
type Profile struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"uniqueIndex;not null" json:"user_id"`
	Name           string    `json:"name"`
	PhoneNumber    string    `json:"phone_number"`
	Address        string    `json:"address"`
	Faculty        string    `json:"faculty"`
	Major          string    `json:"major"`
	EducationLevel string    `json:"education_level"`
	Started        int       `json:"started"`
	Graduated      int       `json:"graduated"`
	CVURL          string    `json:"cv_url"`
	PortfolioURL   string    `json:"portfolio_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ProfileRepository interface {
	Create(ctx context.Context, p *Profile) error
	Update(ctx context.Context, p *Profile) error
	FindByUserID(ctx context.Context, userID uint) (*Profile, error)
}
