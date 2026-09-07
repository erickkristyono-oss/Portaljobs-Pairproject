package handler

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"portaljob/internal/domain"
	"portaljob/internal/usecase"

	"github.com/gin-gonic/gin"
)

// stubProfileRepo is a minimal in-memory ProfileRepository for handler tests.
type stubProfileRepo struct{}

func (s *stubProfileRepo) Create(ctx context.Context, p *domain.Profile) error {
	p.ID = 1
	return nil
}
func (s *stubProfileRepo) Update(ctx context.Context, p *domain.Profile) error { return nil }
func (s *stubProfileRepo) FindByUserID(ctx context.Context, userID uint) (*domain.Profile, error) {
	return nil, domain.ErrNotFound // force the "create" path
}

// saveProfile drives ProfileHandler.Save with a JSON body as an authenticated
// jobseeker (user_id=1) and returns the recorded HTTP response.
func saveProfile(body string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user_id", uint(1)) // same key middleware.Auth sets

	req := httptest.NewRequest(http.MethodPut, "/api/v1/me/profile", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	h := NewProfileHandler(usecase.NewProfileUsecase(&stubProfileRepo{}))
	h.Save(c)
	return w
}

func TestProfileSave_ValidCVURL(t *testing.T) {
	w := saveProfile(`{"name":"Andi","cv_url":"https://drive.google.com/file/d/abc123/view"}`)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d (body: %s)", w.Code, w.Body.String())
	}
}

func TestProfileSave_InvalidCVURL(t *testing.T) {
	w := saveProfile(`{"name":"Andi","cv_url":"bukan-sebuah-url"}`)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for invalid url, got %d (body: %s)", w.Code, w.Body.String())
	}
}

func TestProfileSave_URLOptional(t *testing.T) {
	w := saveProfile(`{"name":"Andi"}`)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 when url omitted, got %d (body: %s)", w.Code, w.Body.String())
	}
}
