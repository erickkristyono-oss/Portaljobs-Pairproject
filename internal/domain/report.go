package domain

import (
	"context"
	"time"
)

const (
	ReportOpen    = "open"
	ReportValid   = "valid"
	ReportInvalid = "invalid"
	ReportClosed  = "closed"
)

// Report is filed by a user against a target (user/company/job).
type Report struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ReporterID uint      `gorm:"index;not null" json:"reporter_id"`
	TargetType string    `gorm:"type:varchar(20)" json:"target_type"`
	TargetID   uint      `json:"target_id"`
	Reason     string    `json:"reason"`
	Status     string    `gorm:"type:varchar(20);default:'open'" json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ReportRepository interface {
	Create(ctx context.Context, r *Report) error
	FindAll(ctx context.Context) ([]Report, error)
	FindByID(ctx context.Context, id uint) (*Report, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
}
