package constant

type ReportTargetType string

const (
	TargetUser    ReportTargetType = "user"
	TargetCompany ReportTargetType = "company"
	TargetJob     ReportTargetType = "job"
)

func IsValidReportTargetType(targetType ReportTargetType) bool {
	switch targetType {
	case TargetUser,
		TargetCompany,
		TargetJob:
		return true
	default:
		return false
	}
}
