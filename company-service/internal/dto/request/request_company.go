package request

type CreateCompanyRequest struct {
	UserID uint `json:"user_id" validate:"required"`
}
