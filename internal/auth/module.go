package auth

import (
	"net/http"
	"time"

	"erp/backend/internal/auth/handler"
	"erp/backend/internal/auth/jwt"
	"erp/backend/internal/auth/middleware"
	"erp/backend/internal/auth/usecase"
	userPostgres "erp/backend/internal/user/repository/postgres"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

const (
	defaultAccessTokenTTL  = 15 * time.Minute
	defaultRefreshTokenTTL = 7 * 24 * time.Hour
	defaultIssuer          = "erp-backend"
)

// Config holds the settings needed to wire the auth module. Secret must be
// loaded from the environment, never hardcoded.
type Config struct {
	Secret          string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	Issuer          string
}

// RegisterModule wires the auth module against the given database (reusing
// the user module's Postgres repository for credential/role lookups) and
// registers its routes. It returns the Authenticate middleware so other
// modules can protect their own routes with it.
func RegisterModule(router *mux.Router, db *gorm.DB, cfg Config) func(http.Handler) http.Handler {
	if cfg.AccessTokenTTL == 0 {
		cfg.AccessTokenTTL = defaultAccessTokenTTL
	}
	if cfg.RefreshTokenTTL == 0 {
		cfg.RefreshTokenTTL = defaultRefreshTokenTTL
	}
	if cfg.Issuer == "" {
		cfg.Issuer = defaultIssuer
	}

	jwtManager := jwt.NewManager(cfg.Secret, cfg.AccessTokenTTL, cfg.RefreshTokenTTL, cfg.Issuer)
	userRepo := userPostgres.NewPostgresUserRepository(db)
	uc := usecase.New(userRepo, jwtManager)
	h := handler.New(uc)
	authenticate := middleware.Authenticate(jwtManager)

	RegisterRoutes(router, h, authenticate)

	return authenticate
}
