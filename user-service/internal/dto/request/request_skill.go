package request

type CreateSkillRequest struct {
	NameLicense  string `json:"name_license" validate:"required"`
	Category     string `json:"category" validate:"required"`
	Subcategory  string `json:"subcategory" validate:"required"`
	Level        string `json:"level" validate:"required,oneof=beginner intermediate expert"`
	Organization string `json:"organization"`
	Grade        string `json:"grade"`
	ExpiredDate  string `json:"expired_date"`
	Description  string `json:"description"`
}

type UpdateSkillRequest struct {
	NameLicense  string `json:"name_license" validate:"required"`
	Category     string `json:"category" validate:"required"`
	Subcategory  string `json:"subcategory" validate:"required"`
	Level        string `json:"level" validate:"required"`
	Organization string `json:"organization"`
	Grade        string `json:"grade"`
	ExpiredDate  string `json:"expired_date"`
	Description  string `json:"description"`
}
