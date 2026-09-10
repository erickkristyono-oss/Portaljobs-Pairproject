package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"job-service/internal/domain/entity"
	"job-service/internal/dto/request"
	"job-service/internal/handler"
	"job-service/internal/usecase/job"

	"github.com/labstack/echo/v5"
)

type mockJobUsecase struct {
	getByIDFunc func(context.Context, uint) (*entity.Job, error)
}

func (m *mockJobUsecase) Create(
	ctx context.Context,
	userID uint,
	req request.CreateJobRequest,
) (*entity.Job, error) {
	return nil, nil
}

func (m *mockJobUsecase) GetByID(
	ctx context.Context,
	id uint,
) (*entity.Job, error) {
	return m.getByIDFunc(ctx, id)
}

func (m *mockJobUsecase) GetAll(
	ctx context.Context,
) ([]*entity.Job, error) {
	return nil, nil
}

func (m *mockJobUsecase) GetByCompanyID(
	ctx context.Context,
	userID uint,
) ([]*entity.Job, error) {
	return nil, nil
}

func (m *mockJobUsecase) GetRequiredSkills(
	ctx context.Context,
	jobID uint,
) ([]*entity.JobRequiredSkill, error) {
	return nil, nil
}

func (m *mockJobUsecase) Update(
	ctx context.Context,
	userID uint,
	id uint,
	req request.UpdateJobRequest,
) (*entity.Job, error) {
	return nil, nil
}

func (m *mockJobUsecase) Delete(
	ctx context.Context,
	userID uint,
	id uint,
) error {
	return nil
}

func (m *mockJobUsecase) Publish(
	ctx context.Context,
	userID uint,
	id uint,
) error {
	return nil
}

func (m *mockJobUsecase) Close(
	ctx context.Context,
	userID uint,
	id uint,
) error {
	return nil
}

var _ job.JobUsecase = (*mockJobUsecase)(nil)

func TestJobHandler_GetByID_Success(t *testing.T) {

	// Mock usecase
	mockUC := &mockJobUsecase{
		getByIDFunc: func(
			ctx context.Context,
			id uint,
		) (*entity.Job, error) {

			return &entity.Job{
				ID:        id,
				CompanyID: 1,
				Judul:     "Backend Developer",
				Lokasi:    "Jakarta",
				Gaji:      10000000,
			}, nil
		},
	}

	// Buat handler menggunakan mock
	h := handler.NewJobHandler(mockUC)

	// Echo
	e := echo.New()

	// Request
	req := httptest.NewRequest(
		http.MethodGet,
		"/jobs/1",
		nil,
	)

	// Response recorder
	rec := httptest.NewRecorder()

	// Context
	c := e.NewContext(req, rec)

	c.SetPath("/jobs/:id")
	c.SetPathValues(echo.PathValues{
		{
			Name:  "id",
			Value: "1",
		},
	})

	// Jalankan handler
	err := h.GetByID(c)

	// Check error
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check status code
	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}
}
