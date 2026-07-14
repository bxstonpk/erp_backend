package usecase

import (
	"erp/backend/internal/auth/dto"

	"golang.org/x/crypto/bcrypt"
)

func (u *authUsecase) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	if user.Status == "inactive" {
		return nil, ErrInvalidCredentials
	}

	roles := roleNames(user.Roles)
	permissions := permissionCodes(user.Permissions)

	pair, err := u.jwtManager.GenerateTokenPair(user.ID, user.Username, roles, permissions)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:           pair.AccessToken,
		RefreshToken:          pair.RefreshToken,
		TokenType:             "Bearer",
		AccessTokenExpiresAt:  pair.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: pair.RefreshTokenExpiresAt,
		User: dto.UserInfo{
			ID:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			FullName:    user.FullName,
			Roles:       roles,
			Permissions: permissions,
		},
	}, nil
}
