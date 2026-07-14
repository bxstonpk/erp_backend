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

func TestHandler_AssignUserRole(t *testing.T) {
	t.Run("success, path userUUID wins over body", func(t *testing.T) {
		var gotReq dto.AssignUserRoleRequest
		uc := &fakeUsecase{
			AssignUserRoleFn: func(req dto.AssignUserRoleRequest) error {
				gotReq = req
				return nil
			},
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.AssignUserRoleRequest{UserUUID: "body-user", RoleUUID: "r-1"})
		rec := httptest.NewRecorder()
		h.AssignUserRole(rec, newRequestWithVars(http.MethodPost, "/users/u-1/roles", body, map[string]string{"userUUID": "u-1"}))

		if rec.Code != http.StatusCreated {
			t.Fatalf("AssignUserRole() status = %d, want %d, body = %s", rec.Code, http.StatusCreated, rec.Body)
		}
		if gotReq.UserUUID != "u-1" || gotReq.RoleUUID != "r-1" {
			t.Fatalf("AssignUserRole() forwarded %+v, want UserUUID=u-1 RoleUUID=r-1", gotReq)
		}
	})

	t.Run("invalid body", func(t *testing.T) {
		h := handler.New(&fakeUsecase{})

		rec := httptest.NewRecorder()
		h.AssignUserRole(rec, newRequestWithVars(http.MethodPost, "/users/u-1/roles", []byte("{"), map[string]string{"userUUID": "u-1"}))

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("AssignUserRole() status = %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})

	t.Run("usecase error", func(t *testing.T) {
		uc := &fakeUsecase{
			AssignUserRoleFn: func(req dto.AssignUserRoleRequest) error { return errors.New("boom") },
		}
		h := handler.New(uc)

		body, _ := json.Marshal(dto.AssignUserRoleRequest{})
		rec := httptest.NewRecorder()
		h.AssignUserRole(rec, newRequestWithVars(http.MethodPost, "/users/u-1/roles", body, map[string]string{"userUUID": "u-1"}))

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("AssignUserRole() status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
	})
}

func TestHandler_AssignUserPermission(t *testing.T) {
	var gotReq dto.AssignUserPermissionRequest
	uc := &fakeUsecase{
		AssignUserPermissionFn: func(req dto.AssignUserPermissionRequest) error {
			gotReq = req
			return nil
		},
	}
	h := handler.New(uc)

	body, _ := json.Marshal(dto.AssignUserPermissionRequest{PermissionUUID: "p-1"})
	rec := httptest.NewRecorder()
	h.AssignUserPermission(rec, newRequestWithVars(http.MethodPost, "/users/u-1/permissions", body, map[string]string{"userUUID": "u-1"}))

	if rec.Code != http.StatusCreated {
		t.Fatalf("AssignUserPermission() status = %d, want %d, body = %s", rec.Code, http.StatusCreated, rec.Body)
	}
	if gotReq.UserUUID != "u-1" || gotReq.PermissionUUID != "p-1" {
		t.Fatalf("AssignUserPermission() forwarded %+v", gotReq)
	}
}

func TestHandler_AssignRolePermission(t *testing.T) {
	var gotReq dto.AssignRolePermissionRequest
	uc := &fakeUsecase{
		AssignRolePermissionFn: func(req dto.AssignRolePermissionRequest) error {
			gotReq = req
			return nil
		},
	}
	h := handler.New(uc)

	body, _ := json.Marshal(dto.AssignRolePermissionRequest{PermissionUUID: "p-1"})
	rec := httptest.NewRecorder()
	h.AssignRolePermission(rec, newRequestWithVars(http.MethodPost, "/roles/r-1/permissions", body, map[string]string{"roleUUID": "r-1"}))

	if rec.Code != http.StatusCreated {
		t.Fatalf("AssignRolePermission() status = %d, want %d, body = %s", rec.Code, http.StatusCreated, rec.Body)
	}
	if gotReq.RoleUUID != "r-1" || gotReq.PermissionUUID != "p-1" {
		t.Fatalf("AssignRolePermission() forwarded %+v", gotReq)
	}
}

func TestHandler_UpdateUserRole(t *testing.T) {
	var gotAssignment string
	var gotReq dto.AssignUserRoleRequest
	uc := &fakeUsecase{
		UpdateUserRoleFn: func(userRoleUUID string, req dto.AssignUserRoleRequest) error {
			gotAssignment, gotReq = userRoleUUID, req
			return nil
		},
	}
	h := handler.New(uc)

	body, _ := json.Marshal(dto.AssignUserRoleRequest{UserUUID: "u-1", RoleUUID: "r-1"})
	rec := httptest.NewRecorder()
	h.UpdateUserRole(rec, newRequestWithVars(http.MethodPut, "/user-roles/ur-1", body, map[string]string{"userRoleUUID": "ur-1"}))

	if rec.Code != http.StatusOK {
		t.Fatalf("UpdateUserRole() status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body)
	}
	if gotAssignment != "ur-1" || gotReq.UserUUID != "u-1" || gotReq.RoleUUID != "r-1" {
		t.Fatalf("UpdateUserRole() forwarded assignment=%q req=%+v", gotAssignment, gotReq)
	}
}

func TestHandler_UpdateUserPermission(t *testing.T) {
	var gotAssignment string
	uc := &fakeUsecase{
		UpdateUserPermissionFn: func(userPermissionUUID string, req dto.AssignUserPermissionRequest) error {
			gotAssignment = userPermissionUUID
			return nil
		},
	}
	h := handler.New(uc)

	body, _ := json.Marshal(dto.AssignUserPermissionRequest{})
	rec := httptest.NewRecorder()
	h.UpdateUserPermission(rec, newRequestWithVars(http.MethodPut, "/user-permissions/up-1", body, map[string]string{"userPermissionUUID": "up-1"}))

	if rec.Code != http.StatusOK {
		t.Fatalf("UpdateUserPermission() status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body)
	}
	if gotAssignment != "up-1" {
		t.Fatalf("UpdateUserPermission() userPermissionUUID = %q, want %q", gotAssignment, "up-1")
	}
}

func TestHandler_UpdateRolePermission(t *testing.T) {
	var gotAssignment string
	uc := &fakeUsecase{
		UpdateRolePermissionFn: func(rolePermissionUUID string, req dto.AssignRolePermissionRequest) error {
			gotAssignment = rolePermissionUUID
			return nil
		},
	}
	h := handler.New(uc)

	body, _ := json.Marshal(dto.AssignRolePermissionRequest{})
	rec := httptest.NewRecorder()
	h.UpdateRolePermission(rec, newRequestWithVars(http.MethodPut, "/role-permissions/rp-1", body, map[string]string{"rolePermissionUUID": "rp-1"}))

	if rec.Code != http.StatusOK {
		t.Fatalf("UpdateRolePermission() status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body)
	}
	if gotAssignment != "rp-1" {
		t.Fatalf("UpdateRolePermission() rolePermissionUUID = %q, want %q", gotAssignment, "rp-1")
	}
}

func TestHandler_DeleteUserRole(t *testing.T) {
	var got string
	uc := &fakeUsecase{
		DeleteUserRoleFn: func(userRoleUUID string) error {
			got = userRoleUUID
			return nil
		},
	}
	h := handler.New(uc)

	rec := httptest.NewRecorder()
	h.DeleteUserRole(rec, newRequestWithVars(http.MethodDelete, "/user-roles/ur-1", nil, map[string]string{"userRoleUUID": "ur-1"}))

	if rec.Code != http.StatusNoContent {
		t.Fatalf("DeleteUserRole() status = %d, want %d", rec.Code, http.StatusNoContent)
	}
	if got != "ur-1" {
		t.Fatalf("DeleteUserRole() userRoleUUID = %q, want %q", got, "ur-1")
	}
}

func TestHandler_DeleteUserPermission(t *testing.T) {
	var got string
	uc := &fakeUsecase{
		DeleteUserPermissionFn: func(userPermissionUUID string) error {
			got = userPermissionUUID
			return nil
		},
	}
	h := handler.New(uc)

	rec := httptest.NewRecorder()
	h.DeleteUserPermission(rec, newRequestWithVars(http.MethodDelete, "/user-permissions/up-1", nil, map[string]string{"userPermissionUUID": "up-1"}))

	if rec.Code != http.StatusNoContent {
		t.Fatalf("DeleteUserPermission() status = %d, want %d", rec.Code, http.StatusNoContent)
	}
	if got != "up-1" {
		t.Fatalf("DeleteUserPermission() userPermissionUUID = %q, want %q", got, "up-1")
	}
}

func TestHandler_DeleteRolePermission(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var got string
		uc := &fakeUsecase{
			DeleteRolePermissionFn: func(rolePermissionUUID string) error {
				got = rolePermissionUUID
				return nil
			},
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.DeleteRolePermission(rec, newRequestWithVars(http.MethodDelete, "/role-permissions/rp-1", nil, map[string]string{"rolePermissionUUID": "rp-1"}))

		if rec.Code != http.StatusNoContent {
			t.Fatalf("DeleteRolePermission() status = %d, want %d", rec.Code, http.StatusNoContent)
		}
		if got != "rp-1" {
			t.Fatalf("DeleteRolePermission() rolePermissionUUID = %q, want %q", got, "rp-1")
		}
	})

	t.Run("usecase error", func(t *testing.T) {
		uc := &fakeUsecase{
			DeleteRolePermissionFn: func(rolePermissionUUID string) error { return errors.New("boom") },
		}
		h := handler.New(uc)

		rec := httptest.NewRecorder()
		h.DeleteRolePermission(rec, newRequestWithVars(http.MethodDelete, "/role-permissions/rp-1", nil, map[string]string{"rolePermissionUUID": "rp-1"}))

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("DeleteRolePermission() status = %d, want %d", rec.Code, http.StatusInternalServerError)
		}
	})
}

func TestHandler_GetUserRoles(t *testing.T) {
	uc := &fakeUsecase{
		GetUserRolesFn: func(userUUID string) ([]*entity.Role, error) {
			if userUUID != "u-1" {
				t.Fatalf("GetUserRoles() userUUID = %q, want %q", userUUID, "u-1")
			}
			return []*entity.Role{{ID: "r-1"}}, nil
		},
	}
	h := handler.New(uc)

	rec := httptest.NewRecorder()
	h.GetUserRoles(rec, newRequestWithVars(http.MethodGet, "/users/u-1/roles", nil, map[string]string{"userUUID": "u-1"}))

	if rec.Code != http.StatusOK {
		t.Fatalf("GetUserRoles() status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body)
	}
}

func TestHandler_GetUserPermissionAssignments(t *testing.T) {
	uc := &fakeUsecase{
		GetUserPermissionAssignmentsFn: func(userUUID string) ([]*entity.PermissionAssignment, error) {
			if userUUID != "u-1" {
				t.Fatalf("GetUserPermissionAssignments() userUUID = %q, want %q", userUUID, "u-1")
			}
			return []*entity.PermissionAssignment{{AssignmentID: "up-1"}}, nil
		},
	}
	h := handler.New(uc)

	rec := httptest.NewRecorder()
	h.GetUserPermissionAssignments(rec, newRequestWithVars(http.MethodGet, "/users/u-1/permissions", nil, map[string]string{"userUUID": "u-1"}))

	if rec.Code != http.StatusOK {
		t.Fatalf("GetUserPermissionAssignments() status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body)
	}
}

func TestHandler_GetRolePermissionAssignments(t *testing.T) {
	uc := &fakeUsecase{
		GetRolePermissionAssignmentsFn: func(roleUUID string) ([]*entity.PermissionAssignment, error) {
			if roleUUID != "r-1" {
				t.Fatalf("GetRolePermissionAssignments() roleUUID = %q, want %q", roleUUID, "r-1")
			}
			return []*entity.PermissionAssignment{{AssignmentID: "rp-1"}}, nil
		},
	}
	h := handler.New(uc)

	rec := httptest.NewRecorder()
	h.GetRolePermissionAssignments(rec, newRequestWithVars(http.MethodGet, "/roles/r-1/permissions", nil, map[string]string{"roleUUID": "r-1"}))

	if rec.Code != http.StatusOK {
		t.Fatalf("GetRolePermissionAssignments() status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body)
	}
}
