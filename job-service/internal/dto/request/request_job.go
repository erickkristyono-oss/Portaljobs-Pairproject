package request

type CreateJobRequest struct {
	Judul            string                          `json:"judul" validate:"required"`
	AboutRole        string                          `json:"about_role" validate:"required"`
	Responsibilities string                          `json:"responsibilities" validate:"required"`
	Deskripsi        string                          `json:"deskripsi" validate:"required"`
	Lokasi           string                          `json:"lokasi" validate:"required"`
	Gaji             int64                           `json:"gaji" validate:"required,min=0"`
	RequiredSkills   []CreateJobRequiredSkillRequest `json:"required_skills"`
}

type UpdateJobRequest struct {
	Judul            string                          `json:"judul" validate:"required"`
	AboutRole        string                          `json:"about_role" validate:"required"`
	Responsibilities string                          `json:"responsibilities" validate:"required"`
	Deskripsi        string                          `json:"deskripsi" validate:"required"`
	Lokasi           string                          `json:"lokasi" validate:"required"`
	Gaji             int64                           `json:"gaji" validate:"required,min=0"`
	RequiredSkills   []CreateJobRequiredSkillRequest `json:"required_skills"`
}
