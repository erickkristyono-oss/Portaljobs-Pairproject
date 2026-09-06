package request

type CreatePortfolioRequest struct {
	NameProject  string `json:"name_project" validate:"required,min=3,max=150"`
	Organization string `json:"organization" validate:"required,max=150"`
	Description  string `json:"description" validate:"required,min=10"`
}

type UpdatePortfolioRequest struct {
	NameProject  string `json:"name_project" validate:"required,min=3,max=150"`
	Organization string `json:"organization" validate:"required,max=150"`
	Description  string `json:"description" validate:"required,min=10"`
}
