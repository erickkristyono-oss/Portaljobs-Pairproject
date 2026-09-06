package entity

import (
	"admin-service/internal/domain/constant"

	"time"
)

type Report struct {
	ID         uint
	ReporterID uint
	TargetType constant.ReportTargetType
	TargetID   uint
	Reason     string
	Status     constant.ReportStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
