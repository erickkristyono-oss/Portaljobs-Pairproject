package entity

import (
	"time"

	"user-service/internal/domain/constant"
)

type User struct {
	ID        uint
	Nama      string
	Email     string
	Password  string
	Role      constant.Role
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
