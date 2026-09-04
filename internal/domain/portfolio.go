package domain

import (
	"context"
	"time"
)

// Portfolio is a project entry owned by a user (1-M)
type Portfolio struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"index;not null" json:"user_id"`
	NameProject  string    `gorm:"not null" json:"name_project"`
	Organization string    `json:"organization"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
}

type PortfolioRepository interface {
	Create(ctx context.Context, p *Portfolio) error
	FindByUserID(ctx context.Context, userID uint) ([]Portfolio, error)
	FindByID(ctx context.Context, id uint) (*Portfolio, error)
	Delete(ctx context.Context, id uint) error
}
