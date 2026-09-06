package dto

type CreateReportRequest struct {
	TargetType string `json:"target_type" validate:"required,oneof=user company job"`
	TargetID   uint   `json:"target_id" validate:"required,min=1"`
	Reason     string `json:"reason" validate:"required,min=10"`
}

type UpdateReportStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=open valid invalid closed"`
}
