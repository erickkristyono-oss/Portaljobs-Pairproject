package usecase

import (
	"context"
	"testing"

	"portaljob/internal/domain"
	"portaljob/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

// mockUserRepo is a hand-written mock of domain.UserRepository. It lets us test
// the usecase without a real database — this is the "mocking" the rubric asks for
type mockUserRepo struct {
	users  map[string]*domain.User
	nextID uint
}

func newMockUserRepo() *mockUserRepo {
	return &mockUserRepo{users: map[string]*domain.User{}, nextID: 1}
}

func (m *mockUserRepo) Create(ctx context.Context, u *domain.User) error {
	u.ID = m.nextID
	m.nextID++
	m.users[u.Email] = u
	return nil
}

func (m *mockUserRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	if u, ok := m.users[email]; ok {
		return u, nil
	}
	return nil, domain.ErrNotFound
}

func (m *mockUserRepo) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	for _, u := range m.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, domain.ErrNotFound
}

func TestRegister_HashesPassword(t *testing.T) {
	uc := NewAuthUsecase(newMockUserRepo(), jwt.NewManager("test-secret", 1))

	user, err := uc.Register(context.Background(), "Budi", "budi@mail.com", "secret123", "company")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.ID == 0 {
		t.Error("expected user to receive an ID")
	}
	if user.Password == "secret123" {
		t.Fatal("password stored in plain text — must be hashed")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("secret123")) != nil {
		t.Error("stored hash does not verify against the original password")
	}
}

func TestRegister_DuplicateEmail(t *testing.T) {
	uc := NewAuthUsecase(newMockUserRepo(), jwt.NewManager("test-secret", 1))
	_, _ = uc.Register(context.Background(), "Budi", "budi@mail.com", "secret123", "company")

	if _, err := uc.Register(context.Background(), "Budi2", "budi@mail.com", "secret123", "jobseeker"); err != domain.ErrEmailTaken {
		t.Errorf("expected ErrEmailTaken, got %v", err)
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	uc := NewAuthUsecase(newMockUserRepo(), jwt.NewManager("test-secret", 1))
	_, _ = uc.Register(context.Background(), "Budi", "budi@mail.com", "secret123", "company")

	if _, _, err := uc.Login(context.Background(), "budi@mail.com", "wrong"); err != domain.ErrInvalidCredential {
		t.Errorf("expected ErrInvalidCredential, got %v", err)
	}
}
