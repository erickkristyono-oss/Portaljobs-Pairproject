package model

import "time"

type UserModel struct {
	ID        uint   `gorm:"primaryKey"`
	Nama      string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"type:varchar(20);not null"`
	Status    string `gorm:"type:varchar(20);default:'active';not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserModel) TableName() string {
	return "users"
}
