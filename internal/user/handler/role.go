package handler

import (
	"net/http"

	"erp/backend/internal/user/dto"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRoleRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	role, err := h.uc.CreateRole(req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusCreated, dto.NewRoleResponse(role))
}

func (h *Handler) GetRole(w http.ResponseWriter, r *http.Request) {
	roleUUID := mux.Vars(r)["roleUUID"]

	role, err := h.uc.GetRole(roleUUID)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewRoleResponse(role))
}

func (h *Handler) ListRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.uc.ListRoles()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewRoleResponses(roles))
}

func (h *Handler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	roleUUID := mux.Vars(r)["roleUUID"]

	var req dto.UpdateRoleRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	role, err := h.uc.UpdateRole(roleUUID, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewRoleResponse(role))
}

func (h *Handler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	roleUUID := mux.Vars(r)["roleUUID"]

	if err := h.uc.DeleteRole(roleUUID); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
