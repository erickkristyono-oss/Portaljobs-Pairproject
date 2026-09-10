package notification

import "context"

type NotificationService interface {
	SendApplicationApplied(ctx context.Context, phone string, jobTitle string) error
	SendInterviewInvitation(ctx context.Context, phone string, jobTitle string, companyName string) error
	SendInterviewSelectedToCompany(ctx context.Context, phone string, candidateName string, jobTitle string) error
	SendApplicationAccepted(ctx context.Context, phone string, jobTitle string, companyName string) error
	SendApplicationAcceptedToCompany(ctx context.Context, phone string, candidateName string, jobTitle string) error
	SendApplicationRejected(ctx context.Context, phone string, jobTitle string) error
}
