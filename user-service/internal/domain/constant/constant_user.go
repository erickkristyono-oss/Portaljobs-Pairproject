package constant

type Role string

const (
	RoleJobseeker Role = "jobseeker"
	RoleCompany   Role = "company"
	RoleAdmin     Role = "admin"
)

func IsValidRole(role Role) bool {
	switch role {
	case RoleJobseeker, RoleCompany, RoleAdmin:
		return true
	default:
		return false
	}
}

const (
	StatusActive    = "active"
	StatusSuspended = "suspended"
)

func IsValidStatus(status string) bool {
	switch status {
	case StatusActive, StatusSuspended:
		return true
	default:
		return false
	}
}
