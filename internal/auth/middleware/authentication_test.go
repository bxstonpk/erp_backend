package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"erp/backend/internal/auth/jwt"
	"erp/backend/internal/auth/middleware"
)

func newTestJWTManager() *jwt.Manager {
	return jwt.NewManager("test-secret", time.Minute, time.Hour, "erp-backend-test")
}

func TestAuthenticate(t *testing.T) {
	jwtManager := newTestJWTManager()

	var gotClaims *jwt.Claims
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotClaims, _ = middleware.ClaimsFromContext(r.Context())
		w.WriteHeader(http.StatusOK)
	})
	handler := middleware.Authenticate(jwtManager)(next)

	t.Run("missing header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusUnauthorized {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
		}
	})

	t.Run("malformed header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
		req.Header.Set("Authorization", "not-bearer")
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusUnauthorized {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
		}
	})

	t.Run("invalid token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
		req.Header.Set("Authorization", "Bearer garbage")
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusUnauthorized {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
		}
	})

	t.Run("valid token sets claims in context", func(t *testing.T) {
		gotClaims = nil
		token, _, err := jwtManager.GenerateAccessToken("u-1", "jdoe", []string{"admin"}, []string{"user:read"})
		if err != nil {
			t.Fatalf("GenerateAccessToken() error = %v", err)
		}

		req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		if gotClaims == nil {
			t.Fatal("claims were not stored in request context")
		}
		if gotClaims.UserID != "u-1" || gotClaims.Username != "jdoe" {
			t.Fatalf("claims = %+v, want UserID=u-1 Username=jdoe", gotClaims)
		}
	})
}
