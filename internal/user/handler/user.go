package handler

import (
	"net/http"

	"erp/backend/internal/user/dto"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.uc.CreateUser(req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusCreated, dto.NewUserResponse(user))
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userUUID := mux.Vars(r)["userUUID"]

	user, err := h.uc.GetUser(userUUID)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewUserResponse(user))
}

func (h *Handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	user, err := h.uc.GetUserByUsername(username)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewUserResponse(user))
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.uc.ListUsers()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewUserResponses(users))
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userUUID := mux.Vars(r)["userUUID"]

	var req dto.UpdateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.uc.UpdateUser(userUUID, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewUserResponse(user))
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userUUID := mux.Vars(r)["userUUID"]

	if err := h.uc.DeleteUser(userUUID); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
