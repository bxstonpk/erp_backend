package dto_test

import (
	"testing"

	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
)

func TestNewPermissionResponse(t *testing.T) {
	permission := &entity.Permission{
		ID: "p-1", CompanyID: "c-1", Code: "user.read", Module: "user", Name: "Read User", Description: "view users",
	}

	got := dto.NewPermissionResponse(permission)

	if got.ID != permission.ID || got.CompanyID != permission.CompanyID || got.Code != permission.Code ||
		got.Module != permission.Module || got.Name != permission.Name || got.Description != permission.Description {
		t.Fatalf("NewPermissionResponse() = %+v, want fields from %+v", got, permission)
	}
}

func TestNewPermissionResponses(t *testing.T) {
	permissions := []*entity.Permission{
		{ID: "p-1", Code: "user.read"},
		{ID: "p-2", Code: "user.write"},
	}

	got := dto.NewPermissionResponses(permissions)

	if len(got) != 2 {
		t.Fatalf("NewPermissionResponses() len = %d, want 2", len(got))
	}
	if got[0].ID != "p-1" || got[1].ID != "p-2" {
		t.Fatalf("NewPermissionResponses() = %+v", got)
	}
}

func TestNewPermissionResponses_empty(t *testing.T) {
	got := dto.NewPermissionResponses(nil)
	if len(got) != 0 {
		t.Fatalf("NewPermissionResponses(nil) len = %d, want 0", len(got))
	}
}
