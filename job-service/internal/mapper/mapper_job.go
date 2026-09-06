package mapper

import (
	"job-service/internal/domain/entity"
	"job-service/internal/dto/response"
	"job-service/internal/repository/postgres/model"
)

func ToJobEntity(jobModel *model.JobModel) *entity.Job {
	if jobModel == nil {
		return nil
	}

	return &entity.Job{
		ID:               jobModel.ID,
		CompanyID:        jobModel.CompanyID,
		Judul:            jobModel.Judul,
		AboutRole:        jobModel.AboutRole,
		Responsibilities: jobModel.Responsibilities,
		Deskripsi:        jobModel.Deskripsi,
		Lokasi:           jobModel.Lokasi,
		Gaji:             jobModel.Gaji,
		Status:           jobModel.Status,
		CreatedAt:        jobModel.CreatedAt,
		UpdatedAt:        jobModel.UpdatedAt,
	}
}

func ToJobModel(job *entity.Job) *model.JobModel {
	if job == nil {
		return nil
	}

	return &model.JobModel{
		ID:               job.ID,
		CompanyID:        job.CompanyID,
		Judul:            job.Judul,
		AboutRole:        job.AboutRole,
		Responsibilities: job.Responsibilities,
		Deskripsi:        job.Deskripsi,
		Lokasi:           job.Lokasi,
		Gaji:             job.Gaji,
		Status:           job.Status,
		CreatedAt:        job.CreatedAt,
		UpdatedAt:        job.UpdatedAt,
	}
}

func ToJobResponse(job *entity.Job) response.JobResponse {
	return response.FromEntity(job)
}
