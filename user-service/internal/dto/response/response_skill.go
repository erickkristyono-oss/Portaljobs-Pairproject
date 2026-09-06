package response

import "time"

type SkillResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	NameLicense  string    `json:"name_license"`
	Level        string    `json:"level"`
	Organization string    `json:"organization"`
	Grade        string    `json:"grade"`
	ExpiredDate  string    `json:"expired_date"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
}
