package constant

type ApplicationStatus string

const (
	ApplicationApplied   ApplicationStatus = "applied"
	ApplicationReviewed  ApplicationStatus = "reviewed"
	ApplicationInterview ApplicationStatus = "interview"
	ApplicationAccepted  ApplicationStatus = "accepted"
	ApplicationRejected  ApplicationStatus = "rejected"
)

func IsValidApplicationStatus(status string) bool {
	switch ApplicationStatus(status) {
	case ApplicationApplied,
		ApplicationReviewed,
		ApplicationInterview,
		ApplicationAccepted,
		ApplicationRejected:
		return true

	default:
		return false
	}
}
