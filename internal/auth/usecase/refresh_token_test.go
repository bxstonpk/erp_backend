package usecase_test

import (
	"errors"
	"testing"
	"time"

	"erp/backend/internal/auth/dto"
	"erp/backend/internal/auth/jwt"
	"erp/backend/internal/auth/usecase"
	"erp/backend/internal/user/entity"
)

func TestAuthUsecase_RefreshToken(t *testing.T) {
	user := &entity.User{ID: "u-1", Username: "jdoe", Roles: []entity.Role{{Name: "admin"}}}

	t.Run("success and rotation", func(t *testing.T) {
		jwtManager := newTestJWTManager()
		repo := &fakeQueryRepository{
			GetUserByUUIDFn: func(userUUID string) (*entity.User, error) {
				if userUUID != "u-1" {
					t.Fatalf("GetUserByUUID() userUUID = %q, want %q", userUUID, "u-1")
				}
				return user, nil
			},
		}
		uc := usecase.New(repo, jwtManager)

		refreshToken, _, err := jwtManager.GenerateRefreshToken("u-1", "jdoe")
		if err != nil {
			t.Fatalf("GenerateRefreshToken() error = %v", err)
		}

		res, err := uc.RefreshToken(dto.RefreshTokenRequest{RefreshToken: refreshToken})
		if err != nil {
			t.Fatalf("RefreshToken() error = %v", err)
		}
		if res.AccessToken == "" || res.RefreshToken == "" {
			t.Fatal("RefreshToken() returned empty tokens")
		}

		// The rotated (old) refresh token must now be rejected as revoked.
		if _, err := uc.RefreshToken(dto.RefreshTokenRequest{RefreshToken: refreshToken}); !errors.Is(err, usecase.ErrTokenRevoked) {
			t.Fatalf("RefreshToken() reused token error = %v, want %v", err, usecase.ErrTokenRevoked)
		}
	})

	t.Run("revoked via logout", func(t *testing.T) {
		jwtManager := newTestJWTManager()
		repo := &fakeQueryRepository{
			GetUserByUUIDFn: func(userUUID string) (*entity.User, error) { return user, nil },
		}
		uc := usecase.New(repo, jwtManager)

		refreshToken, _, err := jwtManager.GenerateRefreshToken("u-1", "jdoe")
		if err != nil {
			t.Fatalf("GenerateRefreshToken() error = %v", err)
		}

		if err := uc.Logout(refreshToken); err != nil {
			t.Fatalf("Logout() error = %v", err)
		}

		if _, err := uc.RefreshToken(dto.RefreshTokenRequest{RefreshToken: refreshToken}); !errors.Is(err, usecase.ErrTokenRevoked) {
			t.Fatalf("RefreshToken() error = %v, want %v", err, usecase.ErrTokenRevoked)
		}
	})

	t.Run("expired token", func(t *testing.T) {
		jwtManager := jwt.NewManager("test-secret", time.Minute, -time.Minute, "erp-backend-test")
		repo := &fakeQueryRepository{}
		uc := usecase.New(repo, jwtManager)

		refreshToken, _, err := jwtManager.GenerateRefreshToken("u-1", "jdoe")
		if err != nil {
			t.Fatalf("GenerateRefreshToken() error = %v", err)
		}

		if _, err := uc.RefreshToken(dto.RefreshTokenRequest{RefreshToken: refreshToken}); !errors.Is(err, jwt.ErrExpiredToken) {
			t.Fatalf("RefreshToken() error = %v, want %v", err, jwt.ErrExpiredToken)
		}
	})

	t.Run("wrong token type", func(t *testing.T) {
		jwtManager := newTestJWTManager()
		repo := &fakeQueryRepository{}
		uc := usecase.New(repo, jwtManager)

		accessToken, _, err := jwtManager.GenerateAccessToken("u-1", "jdoe", nil, nil)
		if err != nil {
			t.Fatalf("GenerateAccessToken() error = %v", err)
		}

		if _, err := uc.RefreshToken(dto.RefreshTokenRequest{RefreshToken: accessToken}); !errors.Is(err, jwt.ErrWrongTokenType) {
			t.Fatalf("RefreshToken() error = %v, want %v", err, jwt.ErrWrongTokenType)
		}
	})
}
