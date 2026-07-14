package usecase

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrTokenRevoked       = errors.New("token has been revoked")
)
