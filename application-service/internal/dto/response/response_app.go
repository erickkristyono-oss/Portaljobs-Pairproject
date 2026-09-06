package response

import (
	"time"

	"application-service/internal/domain/entity"
)

type ApplicationResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	JobID     uint      `json:"job_id"`
	Status    string    `json:"status"`
	AppliedAt time.Time `json:"applied_at"`
}

func FromEntity(application *entity.Application) ApplicationResponse {
	return ApplicationResponse{
		ID:        application.ID,
		UserID:    application.UserID,
		JobID:     application.JobID,
		Status:    application.Status,
		AppliedAt: application.AppliedAt,
	}
}
