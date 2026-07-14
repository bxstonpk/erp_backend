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

func TestHandler_CreatePermission(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		uc := &fakeUsecase{
			CreatePermissionFn: func(req dto.CreatePermissionRequest) (*entity.Permission, error) {
				return &entity.Permission{ID: "p-1", Code: req.Code}, nil
			},
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.CreatePermissionRequest{Code: "user.read"})
		rec := httptest.NewRecorder()
		h.CreatePermission(rec, newRequestWithVars(http.MethodPost, "/permissions", body, nil))

		if rec.Code != http.StatusCreated {
			t.Fatalf("CreatePermission() status = %d, want %d, body = %s", rec.Code, http.StatusCreated, rec.Body)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		h := handler.New(&fakeUsecase{})

		rec := httptest.NewRecorder()
		h.CreatePermission(rec, newRequestWithVars(http.MethodPost, "/permissions", []byte("{"), nil))

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("CreatePermission() status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})

	t.Run("usecase error", func(t *testing.T) {
		uc := &fakeUsecase{
			CreatePermissionFn: func(req dto.CreatePermissionRequest) (*entity.Permission, error) {
				return nil, errors.New("boom")
			},
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.CreatePermissionRequest{})
		rec := httptest.NewRecorder()
		h.CreatePermission(rec, newRequestWithVars(http.MethodPost, "/permissions", body, nil))

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("CreatePermission() status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
	})
}

func TestHandler_GetPermission(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		uc := &fakeUsecase{
			GetPermissionFn: func(permissionUUID string) (*entity.Permission, error) {
				if permissionUUID != "p-1" {
					t.Fatalf("GetPermission() permissionUUID = %q, want %q", permissionUUID, "p-1")
				}
				return &entity.Permission{ID: "p-1"}, nil
			},
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.GetPermission(rec, newRequestWithVars(http.MethodGet, "/permissions/p-1", nil, map[string]string{"permissionUUID": "p-1"}))

		if rec.Code != http.StatusOK {
			t.Fatalf("GetPermission() status = %d, want %d", rec.Code, http.StatusOK)
		}
	})

	t.Run("not found", func(t *testing.T) {
		uc := &fakeUsecase{
			GetPermissionFn: func(permissionUUID string) (*entity.Permission, error) { return nil, errors.New("not found") },
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.GetPermission(rec, newRequestWithVars(http.MethodGet, "/permissions/missing", nil, map[string]string{"permissionUUID": "missing"}))

		if rec.Code != http.StatusNotFound {
			t.Fatalf("GetPermission() status = %d, want %d", rec.Code, http.StatusNotFound)
		}
	})
}

func TestHandler_ListPermissions(t *testing.T) {
	uc := &fakeUsecase{
		ListPermissionsFn: func() ([]*entity.Permission, error) { return []*entity.Permission{{ID: "p-1"}}, nil },
	}
	h := handler.New(uc)

	rec := httptest.NewRecorder()
	h.ListPermissions(rec, newRequestWithVars(http.MethodGet, "/permissions", nil, nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("ListPermissions() status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestHandler_UpdatePermission(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		uc := &fakeUsecase{
			UpdatePermissionFn: func(permissionUUID string, req dto.UpdatePermissionRequest) (*entity.Permission, error) {
				return &entity.Permission{ID: permissionUUID, Code: req.Code}, nil
			},
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.UpdatePermissionRequest{Code: "user.write"})
		rec := httptest.NewRecorder()
		h.UpdatePermission(rec, newRequestWithVars(http.MethodPut, "/permissions/p-1", body, map[string]string{"permissionUUID": "p-1"}))

		if rec.Code != http.StatusOK {
			t.Fatalf("UpdatePermission() status = %d, want %d", rec.Code, http.StatusOK)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		h := handler.New(&fakeUsecase{})

		rec := httptest.NewRecorder()
		h.UpdatePermission(rec, newRequestWithVars(http.MethodPut, "/permissions/p-1", []byte("{"), map[string]string{"permissionUUID": "p-1"}))

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("UpdatePermission() status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})
}

func TestHandler_DeletePermission(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var got string
		uc := &fakeUsecase{
			DeletePermissionFn: func(permissionUUID string) error {
				got = permissionUUID
				return nil
			},
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.DeletePermission(rec, newRequestWithVars(http.MethodDelete, "/permissions/p-1", nil, map[string]string{"permissionUUID": "p-1"}))

		if rec.Code != http.StatusNoContent {
			t.Fatalf("DeletePermission() status = %d, want %d", rec.Code, http.StatusNoContent)
		}
		if got != "p-1" {
			t.Fatalf("DeletePermission() permissionUUID = %q, want %q", got, "p-1")
		}
	})

	t.Run("usecase error", func(t *testing.T) {
		uc := &fakeUsecase{
			DeletePermissionFn: func(permissionUUID string) error { return errors.New("boom") },
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.DeletePermission(rec, newRequestWithVars(http.MethodDelete, "/permissions/p-1", nil, map[string]string{"permissionUUID": "p-1"}))

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("DeletePermission() status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
	})
}
