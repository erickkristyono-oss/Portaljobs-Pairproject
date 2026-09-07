package request

type CreateReportRequest struct {
	TargetType string `json:"target_type" validate:"required"`
	TargetID   uint   `json:"target_id" validate:"required"`
	Reason     string `json:"reason" validate:"required"`
}

type UpdateReportStatusRequest struct {
	Status string `json:"status" validate:"required"`
}
