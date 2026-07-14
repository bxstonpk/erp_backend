package jwt_test

import (
	"errors"
	"testing"
	"time"

	"erp/backend/internal/auth/jwt"
)

func newTestManager() *jwt.Manager {
	return jwt.NewManager("test-secret", time.Minute, time.Hour, "erp-backend-test")
}

func TestManager_GenerateAndValidateAccessToken(t *testing.T) {
	m := newTestManager()

	roles := []string{"admin"}
	permissions := []string{"user:read"}

	token, expiresAt, err := m.GenerateAccessToken("u-1", "jdoe", roles, permissions)
	if err != nil {
		t.Fatalf("GenerateAccessToken() error = %v", err)
	}
	if token == "" {
		t.Fatal("GenerateAccessToken() returned empty token")
	}
	if !expiresAt.After(time.Now()) {
		t.Fatalf("GenerateAccessToken() expiresAt = %v, want future time", expiresAt)
	}

	claims, err := m.ValidateAccessToken(token)
	if err != nil {
		t.Fatalf("ValidateAccessToken() error = %v", err)
	}
	if claims.UserID != "u-1" || claims.Username != "jdoe" {
		t.Fatalf("ValidateAccessToken() claims = %+v, want UserID=u-1 Username=jdoe", claims)
	}
	if len(claims.Roles) != 1 || claims.Roles[0] != "admin" {
		t.Fatalf("ValidateAccessToken() Roles = %v, want [admin]", claims.Roles)
	}
	if len(claims.Permissions) != 1 || claims.Permissions[0] != "user:read" {
		t.Fatalf("ValidateAccessToken() Permissions = %v, want [user:read]", claims.Permissions)
	}
	if claims.Type != jwt.AccessToken {
		t.Fatalf("ValidateAccessToken() Type = %v, want %v", claims.Type, jwt.AccessToken)
	}
}

func TestManager_GenerateTokenPair(t *testing.T) {
	m := newTestManager()

	pair, err := m.GenerateTokenPair("u-1", "jdoe", []string{"admin"}, []string{"user:read"})
	if err != nil {
		t.Fatalf("GenerateTokenPair() error = %v", err)
	}

	if _, err := m.ValidateAccessToken(pair.AccessToken); err != nil {
		t.Fatalf("ValidateAccessToken() on pair.AccessToken error = %v", err)
	}
	if _, err := m.ValidateRefreshToken(pair.RefreshToken); err != nil {
		t.Fatalf("ValidateRefreshToken() on pair.RefreshToken error = %v", err)
	}
}

func TestManager_RejectsWrongTokenType(t *testing.T) {
	m := newTestManager()

	access, _, err := m.GenerateAccessToken("u-1", "jdoe", nil, nil)
	if err != nil {
		t.Fatalf("GenerateAccessToken() error = %v", err)
	}
	if _, err := m.ValidateRefreshToken(access); !errors.Is(err, jwt.ErrWrongTokenType) {
		t.Fatalf("ValidateRefreshToken(access) error = %v, want %v", err, jwt.ErrWrongTokenType)
	}

	refresh, _, err := m.GenerateRefreshToken("u-1", "jdoe")
	if err != nil {
		t.Fatalf("GenerateRefreshToken() error = %v", err)
	}
	if _, err := m.ValidateAccessToken(refresh); !errors.Is(err, jwt.ErrWrongTokenType) {
		t.Fatalf("ValidateAccessToken(refresh) error = %v, want %v", err, jwt.ErrWrongTokenType)
	}
}

func TestManager_RejectsExpiredToken(t *testing.T) {
	m := jwt.NewManager("test-secret", -time.Minute, time.Hour, "erp-backend-test")

	token, _, err := m.GenerateAccessToken("u-1", "jdoe", nil, nil)
	if err != nil {
		t.Fatalf("GenerateAccessToken() error = %v", err)
	}

	if _, err := m.ValidateAccessToken(token); !errors.Is(err, jwt.ErrExpiredToken) {
		t.Fatalf("ValidateAccessToken() error = %v, want %v", err, jwt.ErrExpiredToken)
	}
}

func TestManager_RejectsTokenSignedWithDifferentSecret(t *testing.T) {
	other := jwt.NewManager("a-different-secret", time.Minute, time.Hour, "erp-backend-test")
	token, _, err := other.GenerateAccessToken("u-1", "jdoe", nil, nil)
	if err != nil {
		t.Fatalf("GenerateAccessToken() error = %v", err)
	}

	m := newTestManager()
	if _, err := m.ValidateAccessToken(token); !errors.Is(err, jwt.ErrInvalidToken) {
		t.Fatalf("ValidateAccessToken() error = %v, want %v", err, jwt.ErrInvalidToken)
	}
}

func TestNewManager_PanicsOnEmptySecret(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("NewManager() did not panic on empty secret")
		}
	}()

	jwt.NewManager("", time.Minute, time.Hour, "erp-backend-test")
}
