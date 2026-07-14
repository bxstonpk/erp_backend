package handler

import (
	"net/http"

	"erp/backend/internal/user/dto"

	"github.com/gorilla/mux"
)

func (h *Handler) CreatePermission(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePermissionRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	permission, err := h.uc.CreatePermission(req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusCreated, dto.NewPermissionResponse(permission))
}

func (h *Handler) GetPermission(w http.ResponseWriter, r *http.Request) {
	permissionUUID := mux.Vars(r)["permissionUUID"]

	permission, err := h.uc.GetPermission(permissionUUID)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewPermissionResponse(permission))
}

func (h *Handler) ListPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := h.uc.ListPermissions()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewPermissionResponses(permissions))
}

func (h *Handler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	permissionUUID := mux.Vars(r)["permissionUUID"]

	var req dto.UpdatePermissionRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	permission, err := h.uc.UpdatePermission(permissionUUID, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewPermissionResponse(permission))
}

func (h *Handler) DeletePermission(w http.ResponseWriter, r *http.Request) {
	permissionUUID := mux.Vars(r)["permissionUUID"]

	if err := h.uc.DeletePermission(permissionUUID); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
