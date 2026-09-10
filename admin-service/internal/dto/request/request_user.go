package request

type UpdateUserStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=active suspended"`
}
