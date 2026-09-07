package response

import "time"

type ReportResponse struct {
	ID         uint      `json:"id"`
	ReporterID uint      `json:"reporter_id"`
	TargetType string    `json:"target_type"`
	TargetID   uint      `json:"target_id"`
	Reason     string    `json:"reason"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
