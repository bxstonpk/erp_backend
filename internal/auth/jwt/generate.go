package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenPair struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}

func (m *Manager) newClaims(userID, username string, roles, permissions []string, tokenType TokenType, ttl time.Duration) Claims {
	now := time.Now()

	return Claims{
		UserID:      userID,
		Username:    username,
		Roles:       roles,
		Permissions: permissions,
		Type:        tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			Subject:   userID,
			ID:        uuid.NewString(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}
}

func (m *Manager) sign(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secret)
}

func (m *Manager) GenerateAccessToken(userID, username string, roles, permissions []string) (string, time.Time, error) {
	claims := m.newClaims(userID, username, roles, permissions, AccessToken, m.accessTTL)

	token, err := m.sign(claims)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, claims.ExpiresAt.Time, nil
}

func (m *Manager) GenerateRefreshToken(userID, username string) (string, time.Time, error) {
	claims := m.newClaims(userID, username, nil, nil, RefreshToken, m.refreshTTL)

	token, err := m.sign(claims)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, claims.ExpiresAt.Time, nil
}

func (m *Manager) GenerateTokenPair(userID, username string, roles, permissions []string) (*TokenPair, error) {
	accessToken, accessExpiresAt, err := m.GenerateAccessToken(userID, username, roles, permissions)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExpiresAt, err := m.GenerateRefreshToken(userID, username)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessExpiresAt,
		RefreshTokenExpiresAt: refreshExpiresAt,
	}, nil
}
