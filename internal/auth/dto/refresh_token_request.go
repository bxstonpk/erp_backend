package dto

// RefreshTokenRequest is also used as the Logout request body: logging out
// just means revoking the given refresh token.
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
