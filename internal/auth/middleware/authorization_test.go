package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"erp/backend/internal/auth/middleware"
)

// RequireRole/RequirePermission read claims that Authenticate stores in the
// request context, so these tests chain the real Authenticate middleware in
// front of them rather than injecting claims directly.

func TestRequireRole(t *testing.T) {
	jwtManager := newTestJWTManager()
	authenticate := middleware.Authenticate(jwtManager)

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	handler := authenticate(middleware.RequireRole("admin")(next))

	newRequestWithRoles := func(roles []string) *http.Request {
		token, _, err := jwtManager.GenerateAccessToken("u-1", "jdoe", roles, nil)
		if err != nil {
			t.Fatalf("GenerateAccessToken() error = %v", err)
		}
		req := httptest.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		return req
	}

	t.Run("has required role", func(t *testing.T) {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, newRequestWithRoles([]string{"admin"}))

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
	})

	t.Run("missing required role", func(t *testing.T) {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, newRequestWithRoles([]string{"viewer"}))

		if rec.Code != http.StatusForbidden {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusForbidden)
		}
	})
}

func TestRequirePermission(t *testing.T) {
	jwtManager := newTestJWTManager()
	authenticate := middleware.Authenticate(jwtManager)

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	handler := authenticate(middleware.RequirePermission("user:write")(next))

	newRequestWithPermissions := func(permissions []string) *http.Request {
		token, _, err := jwtManager.GenerateAccessToken("u-1", "jdoe", nil, permissions)
		if err != nil {
			t.Fatalf("GenerateAccessToken() error = %v", err)
		}
		req := httptest.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		return req
	}

	t.Run("has required permission", func(t *testing.T) {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, newRequestWithPermissions([]string{"user:write"}))

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
	})

	t.Run("missing required permission", func(t *testing.T) {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, newRequestWithPermissions([]string{"user:read"}))

		if rec.Code != http.StatusForbidden {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusForbidden)
		}
	})
}
