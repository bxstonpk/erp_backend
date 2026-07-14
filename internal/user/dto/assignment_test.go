package dto_test

import (
	"testing"

	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
)

func TestNewPermissionAssignmentResponse(t *testing.T) {
	assignment := &entity.PermissionAssignment{
		AssignmentID: "up-1",
		Source:       entity.PermissionSourceUser,
		OwnerID:      "u-1",
		Permission:   entity.Permission{ID: "p-1", Code: "user.read"},
	}

	got := dto.NewPermissionAssignmentResponse(assignment)

	if got.AssignmentID != "up-1" || got.Source != "user" || got.OwnerID != "u-1" {
		t.Fatalf("NewPermissionAssignmentResponse() = %+v", got)
	}
	if got.Permission.ID != "p-1" || got.Permission.Code != "user.read" {
		t.Fatalf("NewPermissionAssignmentResponse() Permission = %+v", got.Permission)
	}
}

func TestNewPermissionAssignmentResponse_roleSource(t *testing.T) {
	assignment := &entity.PermissionAssignment{
		AssignmentID: "rp-1",
		Source:       entity.PermissionSourceRole,
		OwnerID:      "r-1",
		Permission:   entity.Permission{ID: "p-1"},
	}

	got := dto.NewPermissionAssignmentResponse(assignment)

	if got.Source != "role" {
		t.Fatalf("NewPermissionAssignmentResponse() Source = %q, want %q", got.Source, "role")
	}
}

func TestNewPermissionAssignmentResponses(t *testing.T) {
	assignments := []*entity.PermissionAssignment{
		{AssignmentID: "up-1", Source: entity.PermissionSourceUser, OwnerID: "u-1"},
		{AssignmentID: "up-2", Source: entity.PermissionSourceUser, OwnerID: "u-1"},
	}

	got := dto.NewPermissionAssignmentResponses(assignments)

	if len(got) != 2 {
		t.Fatalf("NewPermissionAssignmentResponses() len = %d, want 2", len(got))
	}
	if got[0].AssignmentID != "up-1" || got[1].AssignmentID != "up-2" {
		t.Fatalf("NewPermissionAssignmentResponses() = %+v", got)
	}
}

func TestNewPermissionAssignmentResponses_empty(t *testing.T) {
	got := dto.NewPermissionAssignmentResponses(nil)
	if len(got) != 0 {
		t.Fatalf("NewPermissionAssignmentResponses(nil) len = %d, want 0", len(got))
	}
}
