package response

type JobDetailResponse struct {
	ID        uint   `json:"id"`
	CompanyID uint   `json:"company_id"`
	Judul     string `json:"judul"`
	Status    string `json:"status"`
}
