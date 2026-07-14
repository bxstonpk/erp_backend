package handler

import (
	"net/http"

	"erp/backend/internal/user/dto"

	"github.com/gorilla/mux"
)

func (h *Handler) AssignUserCompanyAccess(w http.ResponseWriter, r *http.Request) {
	userUUID := mux.Vars(r)["userUUID"]

	var req dto.CompanyAccessRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.uc.AssignUserCompanyAccess(userUUID, req); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateUserCompanyAccess(w http.ResponseWriter, r *http.Request) {
	userCompanyUUID := mux.Vars(r)["userCompanyUUID"]

	var req dto.UpdateCompanyAccessRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.uc.UpdateUserCompanyAccess(userCompanyUUID, req); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteUserCompanyAccess(w http.ResponseWriter, r *http.Request) {
	userCompanyAccessUUID := mux.Vars(r)["userCompanyUUID"]

	if err := h.uc.DeleteUserCompanyAccess(userCompanyAccessUUID); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetUserCompanyAccesses(w http.ResponseWriter, r *http.Request) {
	userUUID := mux.Vars(r)["userUUID"]

	accesses, err := h.uc.GetUserCompanyAccesses(userUUID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, dto.NewCompanyAccessResponses(accesses))
}
