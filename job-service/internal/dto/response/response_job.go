package response

import (
	"time"

	"job-service/internal/domain/entity"
)

type JobResponse struct {
	ID               uint      `json:"id"`
	CompanyID        uint      `json:"company_id"`
	Judul            string    `json:"judul"`
	AboutRole        string    `json:"about_role"`
	Responsibilities string    `json:"responsibilities"`
	Deskripsi        string    `json:"deskripsi"`
	Lokasi           string    `json:"lokasi"`
	Gaji             int64     `json:"gaji"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func FromEntity(job *entity.Job) JobResponse {
	return JobResponse{
		ID:               job.ID,
		CompanyID:        job.CompanyID,
		Judul:            job.Judul,
		AboutRole:        job.AboutRole,
		Responsibilities: job.Responsibilities,
		Deskripsi:        job.Deskripsi,
		Lokasi:           job.Lokasi,
		Gaji:             job.Gaji,
		Status:           string(job.Status),
		CreatedAt:        job.CreatedAt,
		UpdatedAt:        job.UpdatedAt,
	}
}
