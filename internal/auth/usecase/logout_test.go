package usecase_test

import (
	"errors"
	"testing"

	"erp/backend/internal/auth/jwt"
	"erp/backend/internal/auth/usecase"
)

func TestAuthUsecase_Logout(t *testing.T) {
	t.Run("invalid token", func(t *testing.T) {
		uc := usecase.New(&fakeQueryRepository{}, newTestJWTManager())

		if err := uc.Logout("not-a-jwt"); !errors.Is(err, jwt.ErrInvalidToken) {
			t.Fatalf("Logout() error = %v, want %v", err, jwt.ErrInvalidToken)
		}
	})

	t.Run("access token cannot be logged out as a refresh token", func(t *testing.T) {
		jwtManager := newTestJWTManager()
		uc := usecase.New(&fakeQueryRepository{}, jwtManager)

		accessToken, _, err := jwtManager.GenerateAccessToken("u-1", "jdoe", nil, nil)
		if err != nil {
			t.Fatalf("GenerateAccessToken() error = %v", err)
		}

		if err := uc.Logout(accessToken); !errors.Is(err, jwt.ErrWrongTokenType) {
			t.Fatalf("Logout() error = %v, want %v", err, jwt.ErrWrongTokenType)
		}
	})
}
