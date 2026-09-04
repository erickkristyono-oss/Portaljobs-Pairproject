package handler

type RegisterRequest struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=jobseeker company admin"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateJobRequest struct {
	Judul            string `json:"judul" binding:"required"`
	AboutRole        string `json:"about_role"`
	Responsibilities string `json:"responsibilities"`
	Deskripsi        string `json:"deskripsi"`
	Lokasi           string `json:"lokasi" binding:"required"`
	Gaji             int64  `json:"gaji"`
}

type CompanyRequest struct {
	Name        string `json:"name" binding:"required"`
	FieldOf     string `json:"field_of"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

type ProfileRequest struct {
	Name           string `json:"name" binding:"required"`
	PhoneNumber    string `json:"phone_number"`
	Address        string `json:"address"`
	Faculty        string `json:"faculty"`
	Major          string `json:"major"`
	EducationLevel string `json:"education_level"`
	Started        int    `json:"started"`
	Graduated      int    `json:"graduated"`
}

type SkillRequest struct {
	NameLicense  string `json:"name_license" binding:"required"`
	Level        string `json:"level" binding:"omitempty,oneof=beginner intermediate expert"`
	Organization string `json:"organization"`
	Grade        string `json:"grade"`
	ExpiredDate  string `json:"expired_date"`
	Description  string `json:"description"`
}

type PortfolioRequest struct {
	NameProject  string `json:"name_project" binding:"required"`
	Organization string `json:"organization"`
	Description  string `json:"description"`
}

type DecideRequest struct {
	Status string `json:"status" binding:"required,oneof=interview accepted declined"`
}

type ReportRequest struct {
	TargetType string `json:"target_type" binding:"required,oneof=user company job"`
	TargetID   uint   `json:"target_id" binding:"required"`
	Reason     string `json:"reason" binding:"required"`
}

type ResolveReportRequest struct {
	Valid bool `json:"valid"`
}
