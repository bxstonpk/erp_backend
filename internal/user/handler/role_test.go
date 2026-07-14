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

func TestHandler_CreateRole(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		uc := &fakeUsecase{
			CreateRoleFn: func(req dto.CreateRoleRequest) (*entity.Role, error) {
				return &entity.Role{ID: "r-1", Name: req.Name}, nil
			},
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.CreateRoleRequest{Name: "admin"})
		rec := httptest.NewRecorder()
		h.CreateRole(rec, newRequestWithVars(http.MethodPost, "/roles", body, nil))

		if rec.Code != http.StatusCreated {
			t.Fatalf("CreateRole() status = %d, want %d, body = %s", rec.Code, http.StatusCreated, rec.Body)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		h := handler.New(&fakeUsecase{})

		rec := httptest.NewRecorder()
		h.CreateRole(rec, newRequestWithVars(http.MethodPost, "/roles", []byte("{"), nil))

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("CreateRole() status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})

	t.Run("usecase error", func(t *testing.T) {
		uc := &fakeUsecase{
			CreateRoleFn: func(req dto.CreateRoleRequest) (*entity.Role, error) { return nil, errors.New("boom") },
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.CreateRoleRequest{})
		rec := httptest.NewRecorder()
		h.CreateRole(rec, newRequestWithVars(http.MethodPost, "/roles", body, nil))

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("CreateRole() status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
	})
}

func TestHandler_GetRole(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		uc := &fakeUsecase{
			GetRoleFn: func(roleUUID string) (*entity.Role, error) {
				if roleUUID != "r-1" {
					t.Fatalf("GetRole() roleUUID = %q, want %q", roleUUID, "r-1")
				}
				return &entity.Role{ID: "r-1"}, nil
			},
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.GetRole(rec, newRequestWithVars(http.MethodGet, "/roles/r-1", nil, map[string]string{"roleUUID": "r-1"}))

		if rec.Code != http.StatusOK {
			t.Fatalf("GetRole() status = %d, want %d", rec.Code, http.StatusOK)
		}
	})

	t.Run("not found", func(t *testing.T) {
		uc := &fakeUsecase{
			GetRoleFn: func(roleUUID string) (*entity.Role, error) { return nil, errors.New("not found") },
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.GetRole(rec, newRequestWithVars(http.MethodGet, "/roles/missing", nil, map[string]string{"roleUUID": "missing"}))

		if rec.Code != http.StatusNotFound {
			t.Fatalf("GetRole() status = %d, want %d", rec.Code, http.StatusNotFound)
		}
	})
}

func TestHandler_ListRoles(t *testing.T) {
	uc := &fakeUsecase{
		ListRolesFn: func() ([]*entity.Role, error) { return []*entity.Role{{ID: "r-1"}}, nil },
	}
	h := handler.New(uc)

	rec := httptest.NewRecorder()
	h.ListRoles(rec, newRequestWithVars(http.MethodGet, "/roles", nil, nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("ListRoles() status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestHandler_UpdateRole(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		uc := &fakeUsecase{
			UpdateRoleFn: func(roleUUID string, req dto.UpdateRoleRequest) (*entity.Role, error) {
				return &entity.Role{ID: roleUUID, Name: req.Name}, nil
			},
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.UpdateRoleRequest{Name: "new-name"})
		rec := httptest.NewRecorder()
		h.UpdateRole(rec, newRequestWithVars(http.MethodPut, "/roles/r-1", body, map[string]string{"roleUUID": "r-1"}))

		if rec.Code != http.StatusOK {
			t.Fatalf("UpdateRole() status = %d, want %d", rec.Code, http.StatusOK)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		h := handler.New(&fakeUsecase{})

		rec := httptest.NewRecorder()
		h.UpdateRole(rec, newRequestWithVars(http.MethodPut, "/roles/r-1", []byte("{"), map[string]string{"roleUUID": "r-1"}))

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("UpdateRole() status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})
}

func TestHandler_DeleteRole(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var got string
		uc := &fakeUsecase{
			DeleteRoleFn: func(roleUUID string) error {
				got = roleUUID
				return nil
			},
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.DeleteRole(rec, newRequestWithVars(http.MethodDelete, "/roles/r-1", nil, map[string]string{"roleUUID": "r-1"}))

		if rec.Code != http.StatusNoContent {
			t.Fatalf("DeleteRole() status = %d, want %d", rec.Code, http.StatusNoContent)
		}
		if got != "r-1" {
			t.Fatalf("DeleteRole() roleUUID = %q, want %q", got, "r-1")
		}
	})

	t.Run("usecase error", func(t *testing.T) {
		uc := &fakeUsecase{
			DeleteRoleFn: func(roleUUID string) error { return errors.New("boom") },
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.DeleteRole(rec, newRequestWithVars(http.MethodDelete, "/roles/r-1", nil, map[string]string{"roleUUID": "r-1"}))

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("DeleteRole() status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
	})
}
