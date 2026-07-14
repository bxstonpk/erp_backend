package usecase

import (
	"erp/backend/internal/auth/dto"
	"erp/backend/internal/auth/jwt"
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/repository"
)

type Usecase interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
	Logout(refreshToken string) error
	RefreshToken(req dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
	VerifyToken(accessToken string) (*jwt.Claims, error)
}

type authUsecase struct {
	userRepo   repository.QueryRepository
	jwtManager *jwt.Manager
	revoked    *revokedStore
}

func New(userRepo repository.QueryRepository, jwtManager *jwt.Manager) Usecase {
	return &authUsecase{
		userRepo:   userRepo,
		jwtManager: jwtManager,
		revoked:    newRevokedStore(),
	}
}

func roleNames(roles []entity.Role) []string {
	names := make([]string, 0, len(roles))
	for _, r := range roles {
		names = append(names, r.Name)
	}
	return names
}

func permissionCodes(permissions []entity.Permission) []string {
	codes := make([]string, 0, len(permissions))
	for _, p := range permissions {
		codes = append(codes, p.Code)
	}
	return codes
}
