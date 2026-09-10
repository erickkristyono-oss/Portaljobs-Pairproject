package request

type CreateCompanyProfileRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=150"`
	Phone       string `json:"phone" validate:"required"`
	FieldOf     string `json:"field_of" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateCompanyProfileRequest struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	FieldOf     string `json:"field_of"`
	Address     string `json:"address"`
	Description string `json:"description"`
}
