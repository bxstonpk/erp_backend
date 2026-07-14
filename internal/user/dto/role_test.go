package dto_test

import (
	"testing"

	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
)

func TestNewRoleResponse(t *testing.T) {
	role := &entity.Role{ID: "r-1", CompanyID: "c-1", Name: "admin", Description: "administrator"}

	got := dto.NewRoleResponse(role)

	if got.ID != role.ID || got.CompanyID != role.CompanyID || got.Name != role.Name || got.Description != role.Description {
		t.Fatalf("NewRoleResponse() = %+v, want fields from %+v", got, role)
	}
}

func TestNewRoleResponses(t *testing.T) {
	roles := []*entity.Role{
		{ID: "r-1", Name: "admin"},
		{ID: "r-2", Name: "viewer"},
	}

	got := dto.NewRoleResponses(roles)

	if len(got) != 2 {
		t.Fatalf("NewRoleResponses() len = %d, want 2", len(got))
	}
	if got[0].ID != "r-1" || got[1].ID != "r-2" {
		t.Fatalf("NewRoleResponses() = %+v", got)
	}
}

func TestNewRoleResponses_empty(t *testing.T) {
	got := dto.NewRoleResponses(nil)
	if len(got) != 0 {
		t.Fatalf("NewRoleResponses(nil) len = %d, want 0", len(got))
	}
}
