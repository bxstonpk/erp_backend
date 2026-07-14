package usecase

import "erp/backend/internal/auth/jwt"

func (u *authUsecase) VerifyToken(accessToken string) (*jwt.Claims, error) {
	return u.jwtManager.ValidateAccessToken(accessToken)
}
