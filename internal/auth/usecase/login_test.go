package usecase_test

import (
	"errors"
	"testing"
	"time"

	"erp/backend/internal/auth/dto"
	"erp/backend/internal/auth/jwt"
	"erp/backend/internal/auth/usecase"
	"erp/backend/internal/user/entity"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(t *testing.T, plain string) string {
	t.Helper()

	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("bcrypt.GenerateFromPassword() error = %v", err)
	}

	return string(hashed)
}

func newTestJWTManager() *jwt.Manager {
	return jwt.NewManager("test-secret", time.Minute, time.Hour, "erp-backend-test")
}

func TestAuthUsecase_Login(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		user := &entity.User{
			ID:       "u-1",
			Username: "jdoe",
			Email:    "jdoe@example.com",
			Password: hashPassword(t, "correct-password"),
			FullName: "John Doe",
			Status:   "active",
			Roles:    []entity.Role{{ID: "r-1", Name: "admin"}},
			Permissions: []entity.Permission{
				{ID: "p-1", Code: "user:read"},
			},
		}

		repo := &fakeQueryRepository{
			GetUserByUsernameFn: func(username string) (*entity.User, error) {
				if username != "jdoe" {
					t.Fatalf("GetUserByUsername() username = %q, want %q", username, "jdoe")
				}
				return user, nil
			},
		}
		uc := usecase.New(repo, newTestJWTManager())

		res, err := uc.Login(dto.LoginRequest{Username: "jdoe", Password: "correct-password"})
		if err != nil {
			t.Fatalf("Login() error = %v", err)
		}
		if res.AccessToken == "" || res.RefreshToken == "" {
			t.Fatal("Login() returned empty tokens")
		}
		if res.User.ID != "u-1" || res.User.Username != "jdoe" || res.User.Email != "jdoe@example.com" {
			t.Fatalf("Login() User = %+v, mismatched fields", res.User)
		}
		if len(res.User.Roles) != 1 || res.User.Roles[0] != "admin" {
			t.Fatalf("Login() Roles = %v, want [admin]", res.User.Roles)
		}
		if len(res.User.Permissions) != 1 || res.User.Permissions[0] != "user:read" {
			t.Fatalf("Login() Permissions = %v, want [user:read]", res.User.Permissions)
		}
	})

	t.Run("wrong password", func(t *testing.T) {
		user := &entity.User{
			ID: "u-1", Username: "jdoe", Password: hashPassword(t, "correct-password"), Status: "active",
		}
		repo := &fakeQueryRepository{
			GetUserByUsernameFn: func(username string) (*entity.User, error) { return user, nil },
		}
		uc := usecase.New(repo, newTestJWTManager())

		if _, err := uc.Login(dto.LoginRequest{Username: "jdoe", Password: "wrong-password"}); !errors.Is(err, usecase.ErrInvalidCredentials) {
			t.Fatalf("Login() error = %v, want %v", err, usecase.ErrInvalidCredentials)
		}
	})

	t.Run("unknown user", func(t *testing.T) {
		repo := &fakeQueryRepository{
			GetUserByUsernameFn: func(username string) (*entity.User, error) {
				return nil, errors.New("record not found")
			},
		}
		uc := usecase.New(repo, newTestJWTManager())

		if _, err := uc.Login(dto.LoginRequest{Username: "ghost", Password: "anything"}); !errors.Is(err, usecase.ErrInvalidCredentials) {
			t.Fatalf("Login() error = %v, want %v", err, usecase.ErrInvalidCredentials)
		}
	})

	t.Run("inactive status", func(t *testing.T) {
		user := &entity.User{
			ID: "u-1", Username: "jdoe", Password: hashPassword(t, "correct-password"), Status: "inactive",
		}
		repo := &fakeQueryRepository{
			GetUserByUsernameFn: func(username string) (*entity.User, error) { return user, nil },
		}
		uc := usecase.New(repo, newTestJWTManager())

		if _, err := uc.Login(dto.LoginRequest{Username: "jdoe", Password: "correct-password"}); !errors.Is(err, usecase.ErrInvalidCredentials) {
			t.Fatalf("Login() error = %v, want %v", err, usecase.ErrInvalidCredentials)
		}
	})
}
