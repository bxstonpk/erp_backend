package handler

import "erp/backend/internal/user/usecase"

type Handler struct {
	uc usecase.Usecase
}

func New(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}
