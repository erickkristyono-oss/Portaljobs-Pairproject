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
