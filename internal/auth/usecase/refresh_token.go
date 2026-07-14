package usecase

import "erp/backend/internal/auth/dto"

func (u *authUsecase) RefreshToken(req dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	claims, err := u.jwtManager.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	if u.revoked.isRevoked(claims.ID) {
		return nil, ErrTokenRevoked
	}

	user, err := u.userRepo.GetUserByUUID(claims.UserID)
	if err != nil {
		return nil, err
	}

	// Rotate: the presented refresh token is single-use.
	u.revoked.revoke(claims.ID, claims.ExpiresAt.Time)

	pair, err := u.jwtManager.GenerateTokenPair(user.ID, user.Username, roleNames(user.Roles), permissionCodes(user.Permissions))
	if err != nil {
		return nil, err
	}

	return &dto.RefreshTokenResponse{
		AccessToken:           pair.AccessToken,
		RefreshToken:          pair.RefreshToken,
		TokenType:             "Bearer",
		AccessTokenExpiresAt:  pair.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: pair.RefreshTokenExpiresAt,
	}, nil
}
