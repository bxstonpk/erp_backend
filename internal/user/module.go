package user

import (
	"erp/backend/internal/user/handler"
	"erp/backend/internal/user/repository/postgres"
	"erp/backend/internal/user/usecase"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterModule(router *mux.Router, db *gorm.DB) {
	repo := postgres.NewPostgresUserRepository(db)
	uc := usecase.New(repo)
	h := handler.New(uc)

	RegisterRoutes(router, h)
}
