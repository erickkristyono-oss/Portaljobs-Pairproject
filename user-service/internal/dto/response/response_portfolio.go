package response

import "time"

type PortfolioResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	NameProject  string    `json:"name_project"`
	Organization string    `json:"organization"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
