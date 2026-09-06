package constant

type ApplicationStatus string

const (
	AppSubmitted ApplicationStatus = "submitted"
	AppInterview ApplicationStatus = "interview"
	AppAccepted  ApplicationStatus = "accepted"
	AppDeclined  ApplicationStatus = "declined"
)

func IsValidApplicationStatus(status ApplicationStatus) bool {
	switch status {
	case AppSubmitted,
		AppInterview,
		AppAccepted,
		AppDeclined:
		return true
	default:
		return false
	}
}
