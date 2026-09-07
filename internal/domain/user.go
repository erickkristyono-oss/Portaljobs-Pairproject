package domain

import (
	"context"
	"time"
)

type Role string

const (
	RoleJobseeker Role = "jobseeker"
	RoleCompany   Role = "company"
	RoleAdmin     Role = "admin"
)

func (r Role) Valid() bool {
	switch r {
	case RoleJobseeker, RoleCompany, RoleAdmin:
		return true
	}
	return false
}

const (
	StatusActive    = "active"
	StatusSuspended = "suspended"
)

// User is a pure entity. json:"-" on Password ensures the hash is never
// serialized in any API response.
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"not null" json:"nama"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Role      Role      `gorm:"type:varchar(20);not null" json:"role"`
	Status    string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id uint) (*User, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
	FindAll(ctx context.Context) ([]User, error)
}
