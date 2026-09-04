package domain

import (
	"context"
	"time"
)

const (
	AppSubmitted = "submitted"
	AppInterview = "interview"
	AppAccepted  = "accepted"
	AppDeclined  = "declined"
)

func ValidAppStatus(s string) bool {
	switch s {
	case AppSubmitted, AppInterview, AppAccepted, AppDeclined:
		return true
	}
	return false
}

// Application is a jobseeker's application to a job (1 user - many applications)
type Application struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	JobID     uint      `gorm:"index;not null" json:"job_id"`
	Status    string    `gorm:"type:varchar(20);default:'submitted'" json:"status"`
	AppliedAt time.Time `json:"applied_at"`
}

type ApplicationRepository interface {
	Create(ctx context.Context, a *Application) error
	ExistsByUserAndJob(ctx context.Context, userID, jobID uint) (bool, error)
	FindByUserID(ctx context.Context, userID uint) ([]Application, error)
	FindByJobID(ctx context.Context, jobID uint) ([]Application, error)
	FindByID(ctx context.Context, id uint) (*Application, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
}
