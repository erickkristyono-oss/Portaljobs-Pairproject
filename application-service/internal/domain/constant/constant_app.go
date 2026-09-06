package constant

const (
	ApplicationApplied  = "applied"
	ApplicationReviewed = "reviewed"
	ApplicationAccepted = "accepted"
	ApplicationRejected = "rejected"
)

func IsValidApplicationStatus(status string) bool {
	switch status {
	case ApplicationApplied,
		ApplicationReviewed,
		ApplicationAccepted,
		ApplicationRejected:
		return true
	default:
		return false
	}
}
