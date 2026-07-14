package auth

import (
	"net/http"

	"erp/backend/internal/auth/handler"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, h *handler.Handler, authenticate func(http.Handler) http.Handler) {
	authRoutes := router.PathPrefix("/auth").Subrouter()
	authRoutes.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	authRoutes.HandleFunc("/refresh", h.RefreshToken).Methods(http.MethodPost)
	authRoutes.HandleFunc("/logout", h.Logout).Methods(http.MethodPost)
	authRoutes.Handle("/me", authenticate(http.HandlerFunc(h.Me))).Methods(http.MethodGet)
}
