package model

import "time"

type CompanyModel struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null;uniqueIndex" json:"user_id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	FieldOf     string    `json:"field_of"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (CompanyModel) TableName() string {
	return "companies"
}
