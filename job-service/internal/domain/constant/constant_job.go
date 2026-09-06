package constant

type JobStatus string

const (
	JobDraft     JobStatus = "draft"
	JobPublished JobStatus = "published"
	JobClosed    JobStatus = "closed"
)

func IsValidJobStatus(status JobStatus) bool {
	switch status {
	case JobDraft, JobPublished, JobClosed:
		return true
	default:
		return false
	}
}
