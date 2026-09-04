package usecase

import (
	"context"
	"testing"
	"time"

	"portaljob/internal/domain"
)

// --- hand-written mocks (no DB, no real WA) --

type mockAppRepo struct {
	created  []*domain.Application
	exists   bool
	byID     map[uint]*domain.Application
	statusUp map[uint]string
}

func newMockAppRepo() *mockAppRepo {
	return &mockAppRepo{byID: map[uint]*domain.Application{}, statusUp: map[uint]string{}}
}
func (m *mockAppRepo) Create(ctx context.Context, a *domain.Application) error {
	a.ID = uint(len(m.created) + 1)
	m.created = append(m.created, a)
	m.byID[a.ID] = a
	return nil
}
func (m *mockAppRepo) ExistsByUserAndJob(ctx context.Context, u, j uint) (bool, error) {
	return m.exists, nil
}
func (m *mockAppRepo) FindByUserID(ctx context.Context, u uint) ([]domain.Application, error) {
	return nil, nil
}
func (m *mockAppRepo) FindByJobID(ctx context.Context, j uint) ([]domain.Application, error) {
	return nil, nil
}
func (m *mockAppRepo) FindByID(ctx context.Context, id uint) (*domain.Application, error) {
	if a, ok := m.byID[id]; ok {
		return a, nil
	}
	return nil, domain.ErrNotFound
}
func (m *mockAppRepo) UpdateStatus(ctx context.Context, id uint, s string) error {
	m.statusUp[id] = s
	return nil
}

type mockJobRepo struct{ job *domain.Job }

func (m *mockJobRepo) Create(ctx context.Context, j *domain.Job) error { return nil }
func (m *mockJobRepo) FindAll(ctx context.Context, f domain.JobFilter) ([]domain.Job, error) {
	return nil, nil
}
func (m *mockJobRepo) FindByID(ctx context.Context, id uint) (*domain.Job, error) {
	if m.job != nil && m.job.ID == id {
		return m.job, nil
	}
	return nil, domain.ErrNotFound
}
func (m *mockJobRepo) FindByCompanyID(ctx context.Context, cid uint) ([]domain.Job, error) {
	return nil, nil
}

type mockCompanyRepo struct{ company *domain.Company }

func (m *mockCompanyRepo) Create(ctx context.Context, c *domain.Company) error { return nil }
func (m *mockCompanyRepo) Update(ctx context.Context, c *domain.Company) error { return nil }
func (m *mockCompanyRepo) FindByUserID(ctx context.Context, uid uint) (*domain.Company, error) {
	if m.company != nil && m.company.UserID == uid {
		return m.company, nil
	}
	return nil, domain.ErrNotFound
}
func (m *mockCompanyRepo) FindByID(ctx context.Context, id uint) (*domain.Company, error) {
	return nil, domain.ErrNotFound
}

type mockProfileRepo struct{}

func (m *mockProfileRepo) Create(ctx context.Context, p *domain.Profile) error { return nil }
func (m *mockProfileRepo) Update(ctx context.Context, p *domain.Profile) error { return nil }
func (m *mockProfileRepo) FindByUserID(ctx context.Context, uid uint) (*domain.Profile, error) {
	return nil, domain.ErrNotFound // no phone -> notify is skipped
}

type mockNotifier struct{ sent int }

func (m *mockNotifier) Send(ctx context.Context, to, msg string) error { m.sent++; return nil }

// --- tests ---

func newApplyUC(apps *mockAppRepo, job *domain.Job) *ApplicationUsecase {
	return NewApplicationUsecase(apps, &mockJobRepo{job: job}, &mockCompanyRepo{}, &mockProfileRepo{}, &mockNotifier{})
}

func TestApply_Success(t *testing.T) {
	apps := newMockAppRepo()
	job := &domain.Job{ID: 1, Status: "published"}
	uc := newApplyUC(apps, job)

	app, err := uc.Apply(context.Background(), 10, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if app.Status != domain.AppSubmitted {
		t.Errorf("expected status submitted, got %s", app.Status)
	}
	if len(apps.created) != 1 {
		t.Errorf("expected 1 application created, got %d", len(apps.created))
	}
}

func TestApply_DuplicateReturnsConflict(t *testing.T) {
	apps := newMockAppRepo()
	apps.exists = true
	uc := newApplyUC(apps, &domain.Job{ID: 1, Status: "published"})

	if _, err := uc.Apply(context.Background(), 10, 1); err != domain.ErrConflict {
		t.Errorf("expected ErrConflict, got %v", err)
	}
}

func TestApply_UnpublishedJobRejected(t *testing.T) {
	uc := newApplyUC(newMockAppRepo(), &domain.Job{ID: 1, Status: "closed"})
	if _, err := uc.Apply(context.Background(), 10, 1); err != domain.ErrNotFound {
		t.Errorf("expected ErrNotFound for unpublished job, got %v", err)
	}
}

func TestDecide_NonOwnerForbidden(t *testing.T) {
	apps := newMockAppRepo()
	// application on job #1
	apps.byID[5] = &domain.Application{ID: 5, JobID: 1, UserID: 10, AppliedAt: time.Now()}
	// job #1 belongs to company #99, but the company user (7) owns company #42
	uc := NewApplicationUsecase(
		apps,
		&mockJobRepo{job: &domain.Job{ID: 1, CompanyID: 99, Status: "published"}},
		&mockCompanyRepo{company: &domain.Company{ID: 42, UserID: 7}},
		&mockProfileRepo{},
		&mockNotifier{},
	)
	if _, err := uc.Decide(context.Background(), 7, 5, domain.AppAccepted); err != domain.ErrForbidden {
		t.Errorf("expected ErrForbidden, got %v", err)
	}
}
