package request

type CreateProfileRequest struct {
	Name           string `json:"name" validate:"required,min=3,max=100"`
	PhoneNumber    string `json:"phone_number" validate:"required"`
	Address        string `json:"address" validate:"required"`
	Faculty        string `json:"faculty"`
	Major          string `json:"major"`
	EducationLevel string `json:"education_level" validate:"required"`
	Started        int    `json:"started" validate:"required,min=1900"`
	Graduated      int    `json:"graduated" validate:"required,min=1900"`
}

type UpdateProfileRequest struct {
	Name           string `json:"name" validate:"required,min=3,max=100"`
	PhoneNumber    string `json:"phone_number" validate:"required"`
	Address        string `json:"address" validate:"required"`
	Faculty        string `json:"faculty"`
	Major          string `json:"major"`
	EducationLevel string `json:"education_level" validate:"required"`
	Started        int    `json:"started" validate:"required,min=1900"`
	Graduated      int    `json:"graduated" validate:"required,min=1900"`
}
