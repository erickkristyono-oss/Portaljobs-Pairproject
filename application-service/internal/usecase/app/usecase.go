package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"application-service/internal/client"
	"application-service/internal/domain/constant"
	"application-service/internal/domain/entity"
	domainerrors "application-service/internal/domain/error"
	"application-service/internal/domain/repository"
	"application-service/internal/dto/request"
	"application-service/internal/notification"
	"application-service/internal/skilltag"
)

type applicationUsecase struct {
	applicationRepository repository.ApplicationRepository
	jobClient             *client.JobClient
	userClient            *client.UserClient
	companyClient         *client.CompanyClient
	skillTagUsecase       skilltag.SkillTagUsecase
	notificationService   notification.NotificationService
}

func NewApplicationUsecase(
	applicationRepository repository.ApplicationRepository,
	jobClient *client.JobClient,
	userClient *client.UserClient,
	companyClient *client.CompanyClient,
	skillTagUsecase skilltag.SkillTagUsecase,
	notificationService notification.NotificationService,
) ApplicationUsecase {

	return &applicationUsecase{
		applicationRepository: applicationRepository,
		jobClient:             jobClient,
		userClient:            userClient,
		companyClient:         companyClient,
		skillTagUsecase:       skillTagUsecase,
		notificationService:   notificationService,
	}
}

// Apply digunakan oleh jobseeker untuk melamar pekerjaan.
func (u *applicationUsecase) Apply(ctx context.Context, userID uint, req request.CreateApplicationRequest) (*entity.Application, error) {

	// 1. Cek job
	_, jobStatus, jobTitle, err := u.jobClient.GetJobByID(
		ctx,
		req.JobID,
	)
	if err != nil {
		return nil, err
	}

	// 2. Job harus published
	if jobStatus != "published" {
		return nil, domainerrors.ErrConflict
	}

	// 3. Ambil profile jobseeker
	userProfile, err := u.userClient.GetUserProfile(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	// 4. Ambil skill user
	userSkills, err := u.userClient.GetUserSkills(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	// 5. Ambil required skill dari job
	jobSkills, err := u.jobClient.GetRequiredSkills(
		ctx,
		req.JobID,
	)
	if err != nil {
		return nil, err
	}

	// 6. Cocokkan skill user dengan required skill job
	matchResult := u.skillTagUsecase.Match(
		userSkills,
		jobSkills,
	)

	// 7. Kandidat harus memenuhi skill wajib
	if !matchResult.Eligible {
		return nil, domainerrors.ErrConflict
	}

	// 8. Cek apakah user sudah pernah apply
	existing, err := u.applicationRepository.FindByUserAndJob(
		ctx,
		userID,
		req.JobID,
	)

	if err == nil && existing != nil {
		return nil, domainerrors.ErrConflict
	}

	if err != nil && !errors.Is(
		err,
		domainerrors.ErrNotFound,
	) {
		return nil, err
	}

	// 9. Buat application
	application := &entity.Application{
		UserID:    userID,
		JobID:     req.JobID,
		Status:    string(constant.ApplicationApplied),
		AppliedAt: time.Now(),
	}

	// 10. Simpan ke database
	if err := u.applicationRepository.Create(
		ctx,
		application,
	); err != nil {
		return nil, err
	}

	// 11. Kirim notifikasi WhatsApp ke jobseeker
	if err := u.notificationService.SendApplicationApplied(
		ctx,
		userProfile.PhoneNumber,
		jobTitle,
	); err != nil {
		return nil, err
	}

	return application, nil
}

// GetByID mengambil detail application.
func (u *applicationUsecase) GetByID(ctx context.Context, id uint) (*entity.Application, error) {

	return u.applicationRepository.FindByID(ctx, id)
}

// GetMyApplications mengambil semua lamaran milik jobseeker.
func (u *applicationUsecase) GetMyApplications(ctx context.Context, userID uint) ([]*entity.Application, error) {

	return u.applicationRepository.FindByUserID(
		ctx,
		userID,
	)
}

// GetApplicationsByJobID mengambil semua application
// pada job tertentu.
// Hanya company pemilik job yang boleh mengakses.
func (u *applicationUsecase) GetApplicationsByJobID(ctx context.Context, userID uint, jobID uint) ([]*entity.Application, error) {

	// 1. Ambil informasi job
	companyID, _, _, err := u.jobClient.GetJobByID(
		ctx,
		jobID,
	)
	if err != nil {
		return nil, err
	}

	// 2. Ambil company berdasarkan user yang login
	currentCompanyID, err := u.companyClient.GetCompanyByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	// 3. Pastikan company adalah pemilik job
	if currentCompanyID != companyID {
		return nil, domainerrors.ErrForbidden
	}

	// 4. Ambil applications
	return u.applicationRepository.FindByJobID(
		ctx,
		jobID,
	)
}

func (u *applicationUsecase) UpdateStatus(ctx context.Context, userID uint, id uint, status string) error {

	fmt.Println("========================================")
	fmt.Println("UPDATE APPLICATION STATUS")
	fmt.Println("Application ID :", id)
	fmt.Println("Company User ID :", userID)
	fmt.Println("New Status     :", status)
	fmt.Println("========================================")

	// ========================================
	// STEP 0: VALIDATE STATUS
	// ========================================

	fmt.Println("STEP 0: Validate Status")

	if !constant.IsValidApplicationStatus(status) {
		fmt.Println("ERROR Validate Status:", status)
		return domainerrors.ErrConflict
	}

	fmt.Println("STEP 0 OK")

	// ========================================
	// STEP 1: FIND APPLICATION
	// ========================================

	fmt.Println("STEP 1: FindByID")

	application, err := u.applicationRepository.FindByID(
		ctx,
		id,
	)

	if err != nil {
		fmt.Println("ERROR FindByID:", err)
		return err
	}

	fmt.Println("STEP 1 OK")
	fmt.Println("Application ID     :", application.ID)
	fmt.Println("Application UserID :", application.UserID)
	fmt.Println("Application JobID  :", application.JobID)
	fmt.Println("Application Status :", application.Status)

	// ========================================
	// STEP 2: GET JOB
	// ========================================

	fmt.Println("STEP 2: GetJobByID")

	companyID, jobStatus, jobTitle, err :=
		u.jobClient.GetJobByID(
			ctx,
			application.JobID,
		)

	if err != nil {
		fmt.Println("ERROR GetJobByID:", err)
		return err
	}

	fmt.Println("STEP 2 OK")
	fmt.Println("Company ID :", companyID)
	fmt.Println("Job Status :", jobStatus)
	fmt.Println("Job Title  :", jobTitle)

	// ========================================
	// STEP 3: GET COMPANY BY USER ID
	// ========================================

	fmt.Println("STEP 3: GetCompanyByUserID")

	currentCompanyID, err :=
		u.companyClient.GetCompanyByUserID(
			ctx,
			userID,
		)

	if err != nil {
		fmt.Println("ERROR GetCompanyByUserID:", err)
		return err
	}

	fmt.Println("STEP 3 OK")
	fmt.Println("Current Company ID :", currentCompanyID)

	// ========================================
	// STEP 3.1: CHECK COMPANY OWNERSHIP
	// ========================================

	fmt.Println("STEP 3.1: Check Company Ownership")

	if currentCompanyID != companyID {
		fmt.Println("ERROR: company ID mismatch")
		fmt.Println("Current Company ID :", currentCompanyID)
		fmt.Println("Job Company ID     :", companyID)

		return domainerrors.ErrForbidden
	}

	fmt.Println("STEP 3.1 OK - Company owns this job")

	// ========================================
	// STEP 4: VALIDATE STATUS TRANSITION
	// ========================================

	fmt.Println("STEP 4: Validate Status Transition")

	if !isValidStatusTransition(
		application.Status,
		status,
	) {
		fmt.Println("ERROR: invalid status transition")
		fmt.Println("Current Status :", application.Status)
		fmt.Println("Next Status    :", status)

		return domainerrors.ErrConflict
	}

	fmt.Println("STEP 4 OK")

	// ========================================
	// STEP 5: GET USER PROFILE
	// ========================================

	fmt.Println("STEP 5: GetUserProfile")

	userProfile, err :=
		u.userClient.GetUserProfile(
			ctx,
			application.UserID,
		)

	if err != nil {
		fmt.Println("ERROR GetUserProfile:", err)
		return err
	}

	fmt.Println("STEP 5 OK")
	fmt.Println("Candidate User ID :", application.UserID)
	fmt.Println("Candidate Phone   :", userProfile.PhoneNumber)

	// ========================================
	// STEP 6: GET COMPANY DETAIL
	// ========================================

	fmt.Println("STEP 6: GetCompanyByID")

	company, err :=
		u.companyClient.GetCompanyByID(
			ctx,
			companyID,
		)

	if err != nil {
		fmt.Println("ERROR GetCompanyByID:", err)
		return err
	}

	fmt.Println("STEP 6 OK")
	fmt.Println("Company ID   :", company.ID)
	fmt.Println("Company Name :", company.Name)

	// ========================================
	// STEP 7: UPDATE APPLICATION STATUS
	// ========================================

	fmt.Println("STEP 7: UpdateStatus")

	if err := u.applicationRepository.UpdateStatus(
		ctx,
		id,
		status,
	); err != nil {
		fmt.Println("ERROR UpdateStatus:", err)
		return err
	}

	fmt.Println("STEP 7 OK")
	fmt.Println("Application status successfully updated")
	fmt.Println("New Status :", status)

	// ========================================
	// STEP 8: SEND NOTIFICATION
	// ========================================

	fmt.Println("STEP 8: Send Notification")

	switch constant.ApplicationStatus(status) {

	case constant.ApplicationInterview:

		fmt.Println("Notification: Interview")

		if err := u.notificationService.SendInterviewInvitation(
			ctx,
			userProfile.PhoneNumber,
			jobTitle,
			company.Name,
		); err != nil {

			// Notification error tidak menggagalkan
			// update status application.
			fmt.Println("WARNING SendInterviewInvitation:", err)

		} else {
			fmt.Println("Interview notification sent successfully")
		}

	case constant.ApplicationAccepted:

		fmt.Println("Notification: Accepted")

		if err := u.notificationService.SendApplicationAccepted(
			ctx,
			userProfile.PhoneNumber,
			jobTitle,
			company.Name,
		); err != nil {

			fmt.Println("WARNING SendApplicationAccepted:", err)

		} else {
			fmt.Println("Accepted notification sent successfully")
		}

	case constant.ApplicationRejected:

		fmt.Println("Notification: Rejected")

		if err := u.notificationService.SendApplicationRejected(
			ctx,
			userProfile.PhoneNumber,
			jobTitle,
		); err != nil {

			fmt.Println("WARNING SendApplicationRejected:", err)

		} else {
			fmt.Println("Rejected notification sent successfully")
		}

	default:

		fmt.Println("No notification required for status:", status)
	}

	// ========================================
	// SUCCESS
	// ========================================

	fmt.Println("========================================")
	fmt.Println("UPDATE APPLICATION STATUS SUCCESS")
	fmt.Println("Application ID :", id)
	fmt.Println("Status         :", status)
	fmt.Println("========================================")

	return nil
}

func (u *applicationUsecase) GetCompanyApplications(ctx context.Context, userID uint) ([]*entity.Application, error) {

	// Ambil company ID berdasarkan user yang sedang login
	companyID, err := u.companyClient.GetCompanyByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	// Ambil semua job yang dimiliki company
	jobIDs, err := u.jobClient.GetJobsByCompanyID(
		ctx,
		companyID,
	)
	if err != nil {
		return nil, err
	}

	// Ambil semua application berdasarkan job-job tersebut
	return u.applicationRepository.FindByJobIDs(
		ctx,
		jobIDs,
	)
}

// Delete digunakan untuk menghapus application milik user.
func (u *applicationUsecase) Delete(ctx context.Context, userID uint, id uint) error {

	application, err := u.applicationRepository.FindByID(
		ctx,
		id,
	)
	if err != nil {
		return err
	}

	// Hanya pemilik application yang boleh menghapus.
	if application.UserID != userID {
		return domainerrors.ErrForbidden
	}

	// Application yang sudah diterima/ditolak tidak boleh dihapus.
	if application.Status == string(constant.ApplicationAccepted) ||
		application.Status == string(constant.ApplicationRejected) {
		return domainerrors.ErrConflict
	}

	return u.applicationRepository.Delete(
		ctx,
		id,
	)
}

// Status transition:
//
// applied → reviewed
// reviewed → interview
// interview → accepted
// interview → rejected
func isValidStatusTransition(current string, next string) bool {

	currentStatus := constant.ApplicationStatus(current)
	nextStatus := constant.ApplicationStatus(next)

	switch currentStatus {

	case constant.ApplicationApplied:
		return nextStatus == constant.ApplicationReviewed

	case constant.ApplicationReviewed:
		return nextStatus == constant.ApplicationInterview

	case constant.ApplicationInterview:
		return nextStatus == constant.ApplicationAccepted ||
			nextStatus == constant.ApplicationRejected

	case constant.ApplicationAccepted,
		constant.ApplicationRejected:
		return false

	default:
		return false
	}
}
