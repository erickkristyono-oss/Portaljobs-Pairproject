package response

type JobRequiredSkillResponse struct {
	ID          uint   `json:"id"`
	JobID       uint   `json:"job_id"`
	NameLicense string `json:"name_license"`
	SkillTag    string `json:"skill_tag"`
	Required    bool   `json:"required"`
}
