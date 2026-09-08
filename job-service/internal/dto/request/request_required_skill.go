package request

type CreateJobRequiredSkillRequest struct {
	NameLicense string `json:"name_license" validate:"required"`
	SkillTag    string `json:"skill_tag" validate:"required"`
	Required    bool   `json:"required"`
}
