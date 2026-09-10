package response

type CandidateProfileResponse struct {
	UserID         uint   `json:"user_id"`
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number"`
	Address        string `json:"address"`
	Faculty        string `json:"faculty"`
	Major          string `json:"major"`
	EducationLevel string `json:"education_level"`
	Started        int    `json:"started"`
	Graduated      int    `json:"graduated"`
	CVURL          string `json:"cv_url"`
	PortfolioURL   string `json:"portfolio_url"`
}
