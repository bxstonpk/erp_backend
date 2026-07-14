package jwt

import "time"

// Manager issues and validates the HS256-signed access/refresh token pairs
// used across the API. secret must be non-empty: an empty signing key would
// let anyone forge tokens, so NewManager panics rather than allowing it.
type Manager struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
	issuer     string
}

func NewManager(secret string, accessTTL, refreshTTL time.Duration, issuer string) *Manager {
	if secret == "" {
		panic("jwt: secret must not be empty")
	}

	return &Manager{
		secret:     []byte(secret),
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
		issuer:     issuer,
	}
}
