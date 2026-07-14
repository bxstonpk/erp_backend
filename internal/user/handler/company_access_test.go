package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/handler"
)

func TestHandler_AssignUserCompanyAccess(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var gotUser string
		var gotReq dto.CompanyAccessRequest
		uc := &fakeUsecase{
			AssignUserCompanyAccessFn: func(userUUID string, req dto.CompanyAccessRequest) error {
				gotUser, gotReq = userUUID, req
				return nil
			},
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.CompanyAccessRequest{CompanyID: "c-1"})
		rec := httptest.NewRecorder()
		h.AssignUserCompanyAccess(rec, newRequestWithVars(http.MethodPost, "/users/u-1/company-access", body, map[string]string{"userUUID": "u-1"}))

		if rec.Code != http.StatusCreated {
			t.Fatalf("AssignUserCompanyAccess() status = %d, want %d, body = %s", rec.Code, http.StatusCreated, rec.Body)
		}
		if gotUser != "u-1" || gotReq.CompanyID != "c-1" {
			t.Fatalf("AssignUserCompanyAccess() forwarded user=%q req=%+v", gotUser, gotReq)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		h := handler.New(&fakeUsecase{})

		rec := httptest.NewRecorder()
		h.AssignUserCompanyAccess(rec, newRequestWithVars(http.MethodPost, "/users/u-1/company-access", []byte("{"), map[string]string{"userUUID": "u-1"}))

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("AssignUserCompanyAccess() status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})

	t.Run("usecase error", func(t *testing.T) {
		uc := &fakeUsecase{
			AssignUserCompanyAccessFn: func(userUUID string, req dto.CompanyAccessRequest) error { return errors.New("boom") },
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.CompanyAccessRequest{})
		rec := httptest.NewRecorder()
		h.AssignUserCompanyAccess(rec, newRequestWithVars(http.MethodPost, "/users/u-1/company-access", body, map[string]string{"userUUID": "u-1"}))

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("AssignUserCompanyAccess() status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
	})
}

func TestHandler_UpdateUserCompanyAccess(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var gotUserCompany string
		var gotReq dto.UpdateCompanyAccessRequest
		uc := &fakeUsecase{
			UpdateUserCompanyAccessFn: func(userCompanyUUID string, req dto.UpdateCompanyAccessRequest) error {
				gotUserCompany, gotReq = userCompanyUUID, req
				return nil
			},
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.UpdateCompanyAccessRequest{UserUUID: "u-1", CompanyID: "c-1"})
		rec := httptest.NewRecorder()
		h.UpdateUserCompanyAccess(rec, newRequestWithVars(http.MethodPut, "/user-company-access/uc-1", body, map[string]string{"userCompanyUUID": "uc-1"}))

		if rec.Code != http.StatusOK {
			t.Fatalf("UpdateUserCompanyAccess() status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body)
		}
		if gotUserCompany != "uc-1" || gotReq.UserUUID != "u-1" || gotReq.CompanyID != "c-1" {
			t.Fatalf("UpdateUserCompanyAccess() forwarded userCompany=%q req=%+v", gotUserCompany, gotReq)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		h := handler.New(&fakeUsecase{})

		rec := httptest.NewRecorder()
		h.UpdateUserCompanyAccess(rec, newRequestWithVars(http.MethodPut, "/user-company-access/uc-1", []byte("{"), map[string]string{"userCompanyUUID": "uc-1"}))

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("UpdateUserCompanyAccess() status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})
}

func TestHandler_DeleteUserCompanyAccess(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var got string
		uc := &fakeUsecase{
			DeleteUserCompanyAccessFn: func(userCompanyAccessUUID string) error {
				got = userCompanyAccessUUID
				return nil
			},
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.DeleteUserCompanyAccess(rec, newRequestWithVars(http.MethodDelete, "/user-company-access/uc-1", nil, map[string]string{"userCompanyUUID": "uc-1"}))

		if rec.Code != http.StatusNoContent {
			t.Fatalf("DeleteUserCompanyAccess() status = %d, want %d", rec.Code, http.StatusNoContent)
		}
		if got != "uc-1" {
			t.Fatalf("DeleteUserCompanyAccess() userCompanyAccessUUID = %q, want %q", got, "uc-1")
		}
	})

	t.Run("usecase error", func(t *testing.T) {
		uc := &fakeUsecase{
			DeleteUserCompanyAccessFn: func(userCompanyAccessUUID string) error { return errors.New("boom") },
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.DeleteUserCompanyAccess(rec, newRequestWithVars(http.MethodDelete, "/user-company-access/uc-1", nil, map[string]string{"userCompanyUUID": "uc-1"}))

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("DeleteUserCompanyAccess() status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
	})
}

func TestHandler_GetUserCompanyAccesses(t *testing.T) {
	uc := &fakeUsecase{
		GetUserCompanyAccessesFn: func(userUUID string) ([]*entity.CompanyAccess, error) {
			if userUUID != "u-1" {
				t.Fatalf("GetUserCompanyAccesses() userUUID = %q, want %q", userUUID, "u-1")
			}
			return []*entity.CompanyAccess{{CompanyID: "c-1"}}, nil
		},
	}
	h := handler.New(uc)

	rec := httptest.NewRecorder()
	h.GetUserCompanyAccesses(rec, newRequestWithVars(http.MethodGet, "/users/u-1/company-access", nil, map[string]string{"userUUID": "u-1"}))

	if rec.Code != http.StatusOK {
		t.Fatalf("GetUserCompanyAccesses() status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body)
	}
}
