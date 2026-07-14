package usecase

func (u *authUsecase) Logout(refreshToken string) error {
	claims, err := u.jwtManager.ValidateRefreshToken(refreshToken)
	if err != nil {
		return err
	}

	u.revoked.revoke(claims.ID, claims.ExpiresAt.Time)
	return nil
}
