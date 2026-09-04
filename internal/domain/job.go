package domain

import (
	"context"
	"time"
)

// Job is posted by a company. CompanyID references companies.id
type Job struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	CompanyID        uint      `gorm:"index;not null" json:"company_id"`
	Judul            string    `gorm:"not null" json:"judul"`
	AboutRole        string    `json:"about_role"`
	Responsibilities string    `json:"responsibilities"`
	Deskripsi        string    `json:"deskripsi"`
	Lokasi           string    `json:"lokasi"`
	Gaji             int64     `json:"gaji"`
	Status           string    `gorm:"type:varchar(20);default:'published'" json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type JobFilter struct {
	Lokasi string
	Limit  int
	Offset int
}

type JobRepository interface {
	Create(ctx context.Context, job *Job) error
	FindAll(ctx context.Context, f JobFilter) ([]Job, error)
	FindByID(ctx context.Context, id uint) (*Job, error)
	FindByCompanyID(ctx context.Context, companyID uint) ([]Job, error)
}
