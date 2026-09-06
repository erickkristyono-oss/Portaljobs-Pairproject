package constant

type ReportStatus string

const (
	ReportOpen    ReportStatus = "open"
	ReportValid   ReportStatus = "valid"
	ReportInvalid ReportStatus = "invalid"
	ReportClosed  ReportStatus = "closed"
)

func IsValidReportStatus(status ReportStatus) bool {
	switch status {
	case ReportOpen,
		ReportValid,
		ReportInvalid,
		ReportClosed:
		return true
	default:
		return false
	}
}

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
