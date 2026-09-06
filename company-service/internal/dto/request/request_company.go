package request

type CreateCompanyRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=150"`
	FieldOf     string `json:"field_of" validate:"required,max=100"`
	Address     string `json:"address" validate:"required"`
	Description string `json:"description" validate:"required,min=10"`
}

type UpdateCompanyRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=150"`
	FieldOf     string `json:"field_of" validate:"required,max=100"`
	Address     string `json:"address" validate:"required"`
	Description string `json:"description" validate:"required,min=10"`
}
