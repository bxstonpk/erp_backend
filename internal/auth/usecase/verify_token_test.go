package usecase_test

import (
	"errors"
	"testing"

	"erp/backend/internal/auth/jwt"
	"erp/backend/internal/auth/usecase"
)

func TestAuthUsecase_VerifyToken(t *testing.T) {
	jwtManager := newTestJWTManager()
	uc := usecase.New(&fakeQueryRepository{}, jwtManager)

	t.Run("valid access token", func(t *testing.T) {
		token, _, err := jwtManager.GenerateAccessToken("u-1", "jdoe", []string{"admin"}, nil)
		if err != nil {
			t.Fatalf("GenerateAccessToken() error = %v", err)
		}

		claims, err := uc.VerifyToken(token)
		if err != nil {
			t.Fatalf("VerifyToken() error = %v", err)
		}
		if claims.UserID != "u-1" {
			t.Fatalf("VerifyToken() UserID = %q, want %q", claims.UserID, "u-1")
		}
	})

	t.Run("refresh token rejected", func(t *testing.T) {
		token, _, err := jwtManager.GenerateRefreshToken("u-1", "jdoe")
		if err != nil {
			t.Fatalf("GenerateRefreshToken() error = %v", err)
		}

		if _, err := uc.VerifyToken(token); !errors.Is(err, jwt.ErrWrongTokenType) {
			t.Fatalf("VerifyToken() error = %v, want %v", err, jwt.ErrWrongTokenType)
		}
	})
}
