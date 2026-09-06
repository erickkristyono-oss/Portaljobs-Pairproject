package response

import "time"

type ProfileResponse struct {
	ID             uint      `json:"id"`
	UserID         uint      `json:"user_id"`
	Name           string    `json:"name"`
	PhoneNumber    string    `json:"phone_number"`
	Address        string    `json:"address"`
	Faculty        string    `json:"faculty"`
	Major          string    `json:"major"`
	EducationLevel string    `json:"education_level"`
	Started        int       `json:"started"`
	Graduated      int       `json:"graduated"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
