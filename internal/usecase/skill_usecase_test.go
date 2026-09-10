package usecase

import (
	"context"
	"testing"

	"portaljob/internal/domain"
)

// --- 1. MOCK: struct palsu yang memenuhi domain.SkillRepository ---
// Data disimpan di memori, bukan database
type mockSkillRepo struct {
	byID    map[uint]*domain.Skill
	deleted []uint
}

func newMockSkillRepo() *mockSkillRepo {
	return &mockSkillRepo{byID: map[uint]*domain.Skill{}}
}

func (m *mockSkillRepo) Create(ctx context.Context, s *domain.Skill) error {
	s.ID = uint(len(m.byID) + 1)
	m.byID[s.ID] = s
	return nil
}

func (m *mockSkillRepo) FindByUserID(ctx context.Context, userID uint) ([]domain.Skill, error) {
	var out []domain.Skill
	for _, s := range m.byID {
		if s.UserID == userID {
			out = append(out, *s)
		}
	}
	return out, nil
}

func (m *mockSkillRepo) FindByID(ctx context.Context, id uint) (*domain.Skill, error) {
	if s, ok := m.byID[id]; ok {
		return s, nil
	}
	return nil, domain.ErrNotFound
}

func (m *mockSkillRepo) Delete(ctx context.Context, id uint) error {
	m.deleted = append(m.deleted, id)
	delete(m.byID, id)
	return nil
}

// --- 2 & 3. Suntik mock ke usecase, panggil, lalu periksa hasilnya ---

func TestSkillAdd_Success(t *testing.T) {
	uc := NewSkillUsecase(newMockSkillRepo())
	skill, err := uc.Add(context.Background(), 10, domain.Skill{
		NameLicense: "AWS Certified", Level: domain.LevelIntermediate,
	})
	if err != nil {
		t.Fatalf("tidak menyangka error: %v", err)
	}
	if skill.UserID != 10 {
		t.Errorf("UserID harus 10, dapat %d", skill.UserID)
	}
}

func TestSkillAdd_InvalidLevelRejected(t *testing.T) {
	uc := NewSkillUsecase(newMockSkillRepo())
	_, err := uc.Add(context.Background(), 10, domain.Skill{
		NameLicense: "X", Level: domain.SkillLevel("master"), // bukan beginner/intermediate/expert
	})
	if err != domain.ErrInvalidLevel {
		t.Errorf("harusnya ErrInvalidLevel, dapat %v", err)
	}
}

func TestSkillRemove_NonOwnerForbidden(t *testing.T) {
	repo := newMockSkillRepo()
	// skill #1 milik user 2
	repo.byID[1] = &domain.Skill{ID: 1, UserID: 2, NameLicense: "X"}
	uc := NewSkillUsecase(repo)

	// user 99 mencoba menghapus skill milik user 2
	if err := uc.Remove(context.Background(), 99, 1); err != domain.ErrForbidden {
		t.Errorf("harusnya ErrForbidden, dapat %v", err)
	}
	if len(repo.deleted) != 0 {
		t.Errorf("skill orang lain tidak boleh terhapus")
	}
}

func TestSkillRemove_OwnerSucceeds(t *testing.T) {
	repo := newMockSkillRepo()
	repo.byID[1] = &domain.Skill{ID: 1, UserID: 7, NameLicense: "X"}
	uc := NewSkillUsecase(repo)

	if err := uc.Remove(context.Background(), 7, 1); err != nil {
		t.Fatalf("pemilik harus bisa hapus, dapat error: %v", err)
	}
	if len(repo.deleted) != 1 {
		t.Errorf("skill milik sendiri harusnya terhapus")
	}
}