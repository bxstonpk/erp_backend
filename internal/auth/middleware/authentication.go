package middleware

import (
	"context"
	"net/http"
	"strings"

	"erp/backend/internal/auth/jwt"
)

type contextKey string

const claimsContextKey contextKey = "auth_claims"

// Authenticate returns middleware that requires a valid "Authorization:
// Bearer <token>" access token, storing its claims in the request context
// for downstream handlers and authorization middleware.
func Authenticate(jwtManager *jwt.Manager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, ok := bearerToken(r.Header.Get("Authorization"))
			if !ok {
				http.Error(w, "missing or invalid authorization header", http.StatusUnauthorized)
				return
			}

			claims, err := jwtManager.ValidateAccessToken(token)
			if err != nil {
				http.Error(w, "invalid or expired token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), claimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func bearerToken(header string) (string, bool) {
	const prefix = "Bearer "
	if !strings.HasPrefix(header, prefix) {
		return "", false
	}

	token := strings.TrimSpace(strings.TrimPrefix(header, prefix))
	if token == "" {
		return "", false
	}

	return token, true
}

func ClaimsFromContext(ctx context.Context) (*jwt.Claims, bool) {
	claims, ok := ctx.Value(claimsContextKey).(*jwt.Claims)
	return claims, ok
}
