package request

type CreateApplicationRequest struct {
	JobID uint `json:"job_id" validate:"required"`
}
