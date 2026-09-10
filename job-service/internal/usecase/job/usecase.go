package job

import (
	"context"
	"fmt"

	"job-service/internal/client"
	"job-service/internal/domain/constant"
	"job-service/internal/domain/entity"
	domainerror "job-service/internal/domain/error"
	"job-service/internal/domain/repository"
	"job-service/internal/dto/request"
)

type jobUsecase struct {
	jobRepository           repository.JobRepository
	requiredSkillRepository repository.JobRequiredSkillRepository
	companyClient           *client.CompanyClient
}

var _ JobUsecase = (*jobUsecase)(nil)

func NewJobUsecase(
	jobRepository repository.JobRepository,
	requiredSkillRepository repository.JobRequiredSkillRepository,
	companyClient *client.CompanyClient,
) JobUsecase {
	return &jobUsecase{
		jobRepository:           jobRepository,
		requiredSkillRepository: requiredSkillRepository,
		companyClient:           companyClient,
	}
}

func (u *jobUsecase) Create(
	ctx context.Context,
	userID uint,
	req request.CreateJobRequest,
) (*entity.Job, error) {

	fmt.Println("========================================")
	fmt.Println("JOB USECASE CREATE")
	fmt.Println("User ID:", userID)
	fmt.Println("STEP 1: Get Company ID")
	fmt.Println("========================================")

	companyID, err := u.companyClient.GetCompanyByUserID(
		ctx,
		userID,
	)
	if err != nil {
		fmt.Println("ERROR GET COMPANY:", err)
		return nil, err
	}

	fmt.Println("STEP 1 OK")
	fmt.Println("Company ID:", companyID)

	// ==============================
	// STEP 2: Create Job Entity
	// ==============================
	fmt.Println("STEP 2: Create Job Entity")

	job := &entity.Job{
		CompanyID:        companyID,
		Judul:            req.Judul,
		AboutRole:        req.AboutRole,
		Responsibilities: req.Responsibilities,
		Deskripsi:        req.Deskripsi,
		Lokasi:           req.Lokasi,
		Gaji:             req.Gaji,
		Status:           constant.JobDraft,
	}

	fmt.Println("STEP 2 OK")
	fmt.Println("Job Company ID:", job.CompanyID)

	// ==============================
	// STEP 3: Save Job
	// ==============================
	fmt.Println("STEP 3: Save Job")

	if err := u.jobRepository.Create(
		ctx,
		job,
	); err != nil {
		fmt.Println("ERROR CREATE JOB:", err)
		return nil, err
	}

	fmt.Println("STEP 3 OK")
	fmt.Println("Job ID:", job.ID)

	// ==============================
	// STEP 4: Required Skills
	// ==============================
	fmt.Println("STEP 4: Create Required Skills")

	requiredSkills := make(
		[]entity.JobRequiredSkill,
		0,
		len(req.RequiredSkills),
	)

	for _, skillReq := range req.RequiredSkills {
		requiredSkills = append(
			requiredSkills,
			entity.JobRequiredSkill{
				JobID:       job.ID,
				NameLicense: skillReq.NameLicense,
				SkillTag:    skillReq.SkillTag,
				Required:    skillReq.Required,
			},
		)
	}

	if len(requiredSkills) > 0 {
		if err := u.requiredSkillRepository.CreateMany(
			ctx,
			requiredSkills,
		); err != nil {
			fmt.Println("ERROR CREATE REQUIRED SKILLS:", err)
			return nil, err
		}
	}

	fmt.Println("STEP 4 OK")
	fmt.Println("========================================")
	fmt.Println("JOB CREATE SUCCESS")
	fmt.Println("========================================")

	return job, nil
}

func (u *jobUsecase) GetByID(
	ctx context.Context,
	id uint,
) (*entity.Job, error) {
	return u.jobRepository.FindByID(ctx, id)
}

func (u *jobUsecase) GetAll(
	ctx context.Context,
) ([]*entity.Job, error) {

	// GET /jobs hanya menampilkan job yang published.
	return u.jobRepository.FindPublished(ctx)
}

func (u *jobUsecase) GetByCompanyID(
	ctx context.Context,
	userID uint,
) ([]*entity.Job, error) {

	// JWT memberikan user ID.
	// Cari company ID berdasarkan user ID.
	companyID, err := u.companyClient.GetCompanyByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return u.jobRepository.FindByCompanyID(
		ctx,
		companyID,
	)
}

func (u *jobUsecase) GetRequiredSkills(
	ctx context.Context,
	jobID uint,
) ([]*entity.JobRequiredSkill, error) {

	// Pastikan job tersedia terlebih dahulu.
	_, err := u.jobRepository.FindByID(
		ctx,
		jobID,
	)
	if err != nil {
		return nil, err
	}

	skills, err := u.requiredSkillRepository.FindByJobID(
		ctx,
		jobID,
	)
	if err != nil {
		return nil, err
	}

	result := make(
		[]*entity.JobRequiredSkill,
		0,
		len(skills),
	)

	for i := range skills {
		result = append(
			result,
			&skills[i],
		)
	}

	return result, nil
}

func (u *jobUsecase) Update(
	ctx context.Context,
	userID uint,
	id uint,
	req request.UpdateJobRequest,
) (*entity.Job, error) {

	job, err := u.jobRepository.FindByID(
		ctx,
		id,
	)
	if err != nil {
		return nil, err
	}

	// Ambil company ID berdasarkan user ID dari JWT.
	companyID, err := u.companyClient.GetCompanyByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	// Pastikan job milik company tersebut.
	if job.CompanyID != companyID {
		return nil, domainerror.ErrForbidden
	}

	job.Judul = req.Judul
	job.AboutRole = req.AboutRole
	job.Responsibilities = req.Responsibilities
	job.Deskripsi = req.Deskripsi
	job.Lokasi = req.Lokasi
	job.Gaji = req.Gaji

	if err := u.jobRepository.Update(
		ctx,
		job,
	); err != nil {
		return nil, err
	}

	// Hapus required skills lama.
	if err := u.requiredSkillRepository.DeleteByJobID(
		ctx,
		job.ID,
	); err != nil {
		return nil, err
	}

	// Siapkan required skills baru.
	requiredSkills := make(
		[]entity.JobRequiredSkill,
		0,
		len(req.RequiredSkills),
	)

	for _, skillReq := range req.RequiredSkills {
		requiredSkills = append(
			requiredSkills,
			entity.JobRequiredSkill{
				JobID:       job.ID,
				NameLicense: skillReq.NameLicense,
				SkillTag:    skillReq.SkillTag,
				Required:    skillReq.Required,
			},
		)
	}

	// Simpan required skills baru.
	if err := u.requiredSkillRepository.CreateMany(
		ctx,
		requiredSkills,
	); err != nil {
		return nil, err
	}

	return job, nil
}

func (u *jobUsecase) Delete(
	ctx context.Context,
	userID uint,
	id uint,
) error {

	job, err := u.jobRepository.FindByID(
		ctx,
		id,
	)
	if err != nil {
		return err
	}

	// Ambil company ID berdasarkan user ID dari JWT.
	companyID, err := u.companyClient.GetCompanyByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return err
	}

	// Pastikan job milik company tersebut.
	if job.CompanyID != companyID {
		return domainerror.ErrForbidden
	}

	// Hapus required skills terlebih dahulu.
	if err := u.requiredSkillRepository.DeleteByJobID(
		ctx,
		id,
	); err != nil {
		return err
	}

	return u.jobRepository.Delete(
		ctx,
		id,
	)
}

func (u *jobUsecase) Publish(
	ctx context.Context,
	userID uint,
	id uint,
) error {

	job, err := u.jobRepository.FindByID(
		ctx,
		id,
	)
	if err != nil {
		return err
	}

	// Ambil company ID berdasarkan user ID dari JWT.
	companyID, err := u.companyClient.GetCompanyByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return err
	}

	// Pastikan job milik company tersebut.
	if job.CompanyID != companyID {
		return domainerror.ErrForbidden
	}

	if job.Status != constant.JobDraft {
		return domainerror.ErrConflict
	}

	job.Status = constant.JobPublished

	return u.jobRepository.Update(
		ctx,
		job,
	)
}

func (u *jobUsecase) Close(
	ctx context.Context,
	userID uint,
	id uint,
) error {

	job, err := u.jobRepository.FindByID(
		ctx,
		id,
	)
	if err != nil {
		return err
	}

	// Ambil company ID berdasarkan user ID dari JWT.
	companyID, err := u.companyClient.GetCompanyByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return err
	}

	// Pastikan job milik company tersebut.
	if job.CompanyID != companyID {
		return domainerror.ErrForbidden
	}

	if job.Status != constant.JobPublished {
		return domainerror.ErrConflict
	}

	job.Status = constant.JobClosed

	return u.jobRepository.Update(
		ctx,
		job,
	)
}
