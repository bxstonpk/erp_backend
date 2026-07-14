package dto_test

import (
	"testing"

	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
)

func TestNewUserResponse(t *testing.T) {
	user := &entity.User{
		ID:       "u-1",
		Username: "jdoe",
		Email:    "jdoe@example.com",
		Password: "hashed",
		FullName: "John Doe",
		Phone:    "0800000000",
		Status:   "active",
		Roles:    []entity.Role{{ID: "r-1", Name: "admin"}},
		Permissions: []entity.Permission{
			{ID: "p-1", Code: "user.read"},
		},
		CompanyAccesses: []entity.CompanyAccess{
			{CompanyID: "c-1", BranchID: "b-1", IsDefault: true, Active: true},
		},
	}

	got := dto.NewUserResponse(user)

	if got.ID != user.ID || got.Username != user.Username || got.Email != user.Email ||
		got.FullName != user.FullName || got.Phone != user.Phone || got.Status != user.Status {
		t.Fatalf("NewUserResponse() mismatched scalar fields: %+v", got)
	}
	if len(got.Roles) != 1 || got.Roles[0].ID != "r-1" || got.Roles[0].Name != "admin" {
		t.Fatalf("NewUserResponse() Roles = %+v", got.Roles)
	}
	if len(got.Permissions) != 1 || got.Permissions[0].ID != "p-1" || got.Permissions[0].Code != "user.read" {
		t.Fatalf("NewUserResponse() Permissions = %+v", got.Permissions)
	}
	if len(got.CompanyAccesses) != 1 || got.CompanyAccesses[0].CompanyID != "c-1" || !got.CompanyAccesses[0].IsDefault {
		t.Fatalf("NewUserResponse() CompanyAccesses = %+v", got.CompanyAccesses)
	}
}

func TestNewUserResponse_emptyNestedSlices(t *testing.T) {
	got := dto.NewUserResponse(&entity.User{ID: "u-1"})

	if got.Roles == nil || len(got.Roles) != 0 {
		t.Fatalf("NewUserResponse() Roles = %v, want empty non-nil slice", got.Roles)
	}
	if got.Permissions == nil || len(got.Permissions) != 0 {
		t.Fatalf("NewUserResponse() Permissions = %v, want empty non-nil slice", got.Permissions)
	}
	if got.CompanyAccesses == nil || len(got.CompanyAccesses) != 0 {
		t.Fatalf("NewUserResponse() CompanyAccesses = %v, want empty non-nil slice", got.CompanyAccesses)
	}
}

func TestNewUserResponses(t *testing.T) {
	users := []*entity.User{
		{ID: "u-1", Username: "a"},
		{ID: "u-2", Username: "b"},
	}

	got := dto.NewUserResponses(users)

	if len(got) != 2 {
		t.Fatalf("NewUserResponses() len = %d, want 2", len(got))
	}
	if got[0].ID != "u-1" || got[1].ID != "u-2" {
		t.Fatalf("NewUserResponses() = %+v", got)
	}
}

func TestNewUserResponses_empty(t *testing.T) {
	got := dto.NewUserResponses(nil)
	if len(got) != 0 {
		t.Fatalf("NewUserResponses(nil) len = %d, want 0", len(got))
	}
}
