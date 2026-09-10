package response

type CompanyDetailResponse struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
}
