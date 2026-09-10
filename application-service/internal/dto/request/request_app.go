package request

type CreateApplicationRequest struct {
	JobID uint `json:"job_id" validate:"required"`
}

type UpdateApplicationStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=applied reviewed interview accepted rejected"`
}
