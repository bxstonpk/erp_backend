package handler

import (
	"net/http"

	"erp/backend/internal/user/dto"

	"github.com/gorilla/mux"
)

func (h *Handler) AssignUserRole(w http.ResponseWriter, r *http.Request) {
	var req dto.AssignUserRoleRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	req.UserUUID = mux.Vars(r)["userUUID"]

	if err := h.uc.AssignUserRole(req); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) AssignUserPermission(w http.ResponseWriter, r *http.Request) {
	var req dto.AssignUserPermissionRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	req.UserUUID = mux.Vars(r)["userUUID"]

	if err := h.uc.AssignUserPermission(req); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) AssignRolePermission(w http.ResponseWriter, r *http.Request) {
	var req dto.AssignRolePermissionRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	req.RoleUUID = mux.Vars(r)["roleUUID"]

	if err := h.uc.AssignRolePermission(req); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	userRoleUUID := mux.Vars(r)["userRoleUUID"]

	var req dto.AssignUserRoleRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.uc.UpdateUserRole(userRoleUUID, req); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateUserPermission(w http.ResponseWriter, r *http.Request) {
	userPermissionUUID := mux.Vars(r)["userPermissionUUID"]

	var req dto.AssignUserPermissionRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.uc.UpdateUserPermission(userPermissionUUID, req); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateRolePermission(w http.ResponseWriter, r *http.Request) {
	rolePermissionUUID := mux.Vars(r)["rolePermissionUUID"]

	var req dto.AssignRolePermissionRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.uc.UpdateRolePermission(rolePermissionUUID, req); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteUserRole(w http.ResponseWriter, r *http.Request) {
	userRoleUUID := mux.Vars(r)["userRoleUUID"]

	if err := h.uc.DeleteUserRole(userRoleUUID); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteUserPermission(w http.ResponseWriter, r *http.Request) {
	userPermissionUUID := mux.Vars(r)["userPermissionUUID"]

	if err := h.uc.DeleteUserPermission(userPermissionUUID); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteRolePermission(w http.ResponseWriter, r *http.Request) {
	rolePermissionUUID := mux.Vars(r)["rolePermissionUUID"]

	if err := h.uc.DeleteRolePermission(rolePermissionUUID); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetUserRoles(w http.ResponseWriter, r *http.Request) {
	userUUID := mux.Vars(r)["userUUID"]

	roles, err := h.uc.GetUserRoles(userUUID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewRoleResponses(roles))
}

func (h *Handler) GetUserPermissionAssignments(w http.ResponseWriter, r *http.Request) {
	userUUID := mux.Vars(r)["userUUID"]

	assignments, err := h.uc.GetUserPermissionAssignments(userUUID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewPermissionAssignmentResponses(assignments))
}

func (h *Handler) GetRolePermissionAssignments(w http.ResponseWriter, r *http.Request) {
	roleUUID := mux.Vars(r)["roleUUID"]

	assignments, err := h.uc.GetRolePermissionAssignments(roleUUID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewPermissionAssignmentResponses(assignments))
}
