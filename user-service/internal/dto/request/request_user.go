package request

type RegisterRequest struct {
	Nama     string `json:"nama" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,oneof=jobseeker company admin"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=active suspended"`
}
