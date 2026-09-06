package entity

import (
	"job-service/internal/domain/constant"
	"time"
)

type Job struct {
	ID               uint
	CompanyID        uint
	Judul            string
	AboutRole        string
	Responsibilities string
	Deskripsi        string
	Lokasi           string
	Gaji             int64
	Status           constant.JobStatus
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
