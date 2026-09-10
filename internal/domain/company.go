package domain

import (
	"context"
	"time"
)

// Company is the company profile, linked 1-1 to a user whose role is "company".
type Company struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"uniqueIndex;not null" json:"user_id"`
	Name        string    `gorm:"not null" json:"name"`
	FieldOf     string    `json:"field_of"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CompanyRepository interface {
	Create(ctx context.Context, c *Company) error
	Update(ctx context.Context, c *Company) error
	FindByUserID(ctx context.Context, userID uint) (*Company, error)
	FindByID(ctx context.Context, id uint) (*Company, error)
}
