package usecase_test

import (
	"errors"
	"testing"

	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/usecase"
)

func TestAssignmentUsecase_AssignUserRole(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var gotUser, gotRole string
		repo := &fakeRepository{
			AssignUserRoleFn: func(userUUID, roleUUID string) error {
				gotUser, gotRole = userUUID, roleUUID
				return nil
			},
		}
		uc := usecase.New(repo)

		req := dto.AssignUserRoleRequest{UserUUID: "u-1", RoleUUID: "r-1"}
		if err := uc.AssignUserRole(req); err != nil {
			t.Fatalf("AssignUserRole() error = %v", err)
		}
		if gotUser != "u-1" || gotRole != "r-1" {
			t.Fatalf("AssignUserRole() forwarded (%q, %q), want (%q, %q)", gotUser, gotRole, "u-1", "r-1")
		}
	})

	t.Run("repository error", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			AssignUserRoleFn: func(userUUID, roleUUID string) error { return wantErr },
		}
		uc := usecase.New(repo)

		if err := uc.AssignUserRole(dto.AssignUserRoleRequest{}); !errors.Is(err, wantErr) {
			t.Fatalf("AssignUserRole() error = %v, want %v", err, wantErr)
		}
	})
}

func TestAssignmentUsecase_AssignUserPermission(t *testing.T) {
	var gotUser, gotPermission string
	repo := &fakeRepository{
		AssignUserPermissionFn: func(userUUID, permissionUUID string) error {
			gotUser, gotPermission = userUUID, permissionUUID
			return nil
		},
	}
	uc := usecase.New(repo)

	req := dto.AssignUserPermissionRequest{UserUUID: "u-1", PermissionUUID: "p-1"}
	if err := uc.AssignUserPermission(req); err != nil {
		t.Fatalf("AssignUserPermission() error = %v", err)
	}
	if gotUser != "u-1" || gotPermission != "p-1" {
		t.Fatalf("AssignUserPermission() forwarded (%q, %q), want (%q, %q)", gotUser, gotPermission, "u-1", "p-1")
	}
}

func TestAssignmentUsecase_AssignRolePermission(t *testing.T) {
	var gotRole, gotPermission string
	repo := &fakeRepository{
		AssignRolePermissionFn: func(roleUUID, permissionUUID string) error {
			gotRole, gotPermission = roleUUID, permissionUUID
			return nil
		},
	}
	uc := usecase.New(repo)

	req := dto.AssignRolePermissionRequest{RoleUUID: "r-1", PermissionUUID: "p-1"}
	if err := uc.AssignRolePermission(req); err != nil {
		t.Fatalf("AssignRolePermission() error = %v", err)
	}
	if gotRole != "r-1" || gotPermission != "p-1" {
		t.Fatalf("AssignRolePermission() forwarded (%q, %q), want (%q, %q)", gotRole, gotPermission, "r-1", "p-1")
	}
}

func TestAssignmentUsecase_UpdateUserRole(t *testing.T) {
	var gotAssignment, gotUser, gotRole string
	repo := &fakeRepository{
		UpdateUserRoleFn: func(userRoleUUID, userUUID, roleUUID string) error {
			gotAssignment, gotUser, gotRole = userRoleUUID, userUUID, roleUUID
			return nil
		},
	}
	uc := usecase.New(repo)

	req := dto.AssignUserRoleRequest{UserUUID: "u-1", RoleUUID: "r-1"}
	if err := uc.UpdateUserRole("ur-1", req); err != nil {
		t.Fatalf("UpdateUserRole() error = %v", err)
	}
	if gotAssignment != "ur-1" || gotUser != "u-1" || gotRole != "r-1" {
		t.Fatalf("UpdateUserRole() forwarded (%q, %q, %q), want (%q, %q, %q)",
			gotAssignment, gotUser, gotRole, "ur-1", "u-1", "r-1")
	}
}

func TestAssignmentUsecase_UpdateUserPermission(t *testing.T) {
	var gotAssignment, gotUser, gotPermission string
	repo := &fakeRepository{
		UpdateUserPermissionFn: func(userPermissionUUID, userUUID, permissionUUID string) error {
			gotAssignment, gotUser, gotPermission = userPermissionUUID, userUUID, permissionUUID
			return nil
		},
	}
	uc := usecase.New(repo)

	req := dto.AssignUserPermissionRequest{UserUUID: "u-1", PermissionUUID: "p-1"}
	if err := uc.UpdateUserPermission("up-1", req); err != nil {
		t.Fatalf("UpdateUserPermission() error = %v", err)
	}
	if gotAssignment != "up-1" || gotUser != "u-1" || gotPermission != "p-1" {
		t.Fatalf("UpdateUserPermission() forwarded (%q, %q, %q), want (%q, %q, %q)",
			gotAssignment, gotUser, gotPermission, "up-1", "u-1", "p-1")
	}
}

func TestAssignmentUsecase_UpdateRolePermission(t *testing.T) {
	var gotAssignment, gotRole, gotPermission string
	repo := &fakeRepository{
		UpdateRolePermissionFn: func(rolePermissionUUID, roleUUID, permissionUUID string) error {
			gotAssignment, gotRole, gotPermission = rolePermissionUUID, roleUUID, permissionUUID
			return nil
		},
	}
	uc := usecase.New(repo)

	req := dto.AssignRolePermissionRequest{RoleUUID: "r-1", PermissionUUID: "p-1"}
	if err := uc.UpdateRolePermission("rp-1", req); err != nil {
		t.Fatalf("UpdateRolePermission() error = %v", err)
	}
	if gotAssignment != "rp-1" || gotRole != "r-1" || gotPermission != "p-1" {
		t.Fatalf("UpdateRolePermission() forwarded (%q, %q, %q), want (%q, %q, %q)",
			gotAssignment, gotRole, gotPermission, "rp-1", "r-1", "p-1")
	}
}

func TestAssignmentUsecase_Delete(t *testing.T) {
	t.Run("DeleteUserRole", func(t *testing.T) {
		var got string
		repo := &fakeRepository{
			DeleteUserRoleFn: func(userRoleUUID string) error {
				got = userRoleUUID
				return nil
			},
		}
		uc := usecase.New(repo)

		if err := uc.DeleteUserRole("ur-1"); err != nil {
			t.Fatalf("DeleteUserRole() error = %v", err)
		}
		if got != "ur-1" {
			t.Fatalf("DeleteUserRole() userRoleUUID = %q, want %q", got, "ur-1")
		}
	})

	t.Run("DeleteUserPermission", func(t *testing.T) {
		var got string
		repo := &fakeRepository{
			DeleteUserPermissionFn: func(userPermissionUUID string) error {
				got = userPermissionUUID
				return nil
			},
		}
		uc := usecase.New(repo)

		if err := uc.DeleteUserPermission("up-1"); err != nil {
			t.Fatalf("DeleteUserPermission() error = %v", err)
		}
		if got != "up-1" {
			t.Fatalf("DeleteUserPermission() userPermissionUUID = %q, want %q", got, "up-1")
		}
	})

	t.Run("DeleteRolePermission", func(t *testing.T) {
		var got string
		repo := &fakeRepository{
			DeleteRolePermissionFn: func(rolePermissionUUID string) error {
				got = rolePermissionUUID
				return nil
			},
		}
		uc := usecase.New(repo)

		if err := uc.DeleteRolePermission("rp-1"); err != nil {
			t.Fatalf("DeleteRolePermission() error = %v", err)
		}
		if got != "rp-1" {
			t.Fatalf("DeleteRolePermission() rolePermissionUUID = %q, want %q", got, "rp-1")
		}
	})

	t.Run("repository error propagates", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			DeleteUserRoleFn: func(userRoleUUID string) error { return wantErr },
		}
		uc := usecase.New(repo)

		if err := uc.DeleteUserRole("ur-1"); !errors.Is(err, wantErr) {
			t.Fatalf("DeleteUserRole() error = %v, want %v", err, wantErr)
		}
	})
}

func TestAssignmentUsecase_GetUserRoles(t *testing.T) {
	want := []*entity.Role{{ID: "r-1"}, {ID: "r-2"}}
	repo := &fakeRepository{
		GetUserRolesFn: func(userUUID string) ([]*entity.Role, error) {
			if userUUID != "u-1" {
				t.Fatalf("GetUserRoles() userUUID = %q, want %q", userUUID, "u-1")
			}
			return want, nil
		},
	}
	uc := usecase.New(repo)

	got, err := uc.GetUserRoles("u-1")
	if err != nil {
		t.Fatalf("GetUserRoles() error = %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("GetUserRoles() len = %d, want %d", len(got), len(want))
	}
}

func TestAssignmentUsecase_GetUserPermissionAssignments(t *testing.T) {
	want := []*entity.PermissionAssignment{{AssignmentID: "up-1", Source: entity.PermissionSourceUser, OwnerID: "u-1"}}
	repo := &fakeRepository{
		GetPermissionAssignmentsByUserFn: func(userUUID string) ([]*entity.PermissionAssignment, error) {
			if userUUID != "u-1" {
				t.Fatalf("GetPermissionAssignmentsByUser() userUUID = %q, want %q", userUUID, "u-1")
			}
			return want, nil
		},
	}
	uc := usecase.New(repo)

	got, err := uc.GetUserPermissionAssignments("u-1")
	if err != nil {
		t.Fatalf("GetUserPermissionAssignments() error = %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("GetUserPermissionAssignments() len = %d, want %d", len(got), len(want))
	}
}

func TestAssignmentUsecase_GetRolePermissionAssignments(t *testing.T) {
	want := []*entity.PermissionAssignment{{AssignmentID: "rp-1", Source: entity.PermissionSourceRole, OwnerID: "r-1"}}
	repo := &fakeRepository{
		GetPermissionAssignmentsByRoleFn: func(roleUUID string) ([]*entity.PermissionAssignment, error) {
			if roleUUID != "r-1" {
				t.Fatalf("GetPermissionAssignmentsByRole() roleUUID = %q, want %q", roleUUID, "r-1")
			}
			return want, nil
		},
	}
	uc := usecase.New(repo)

	got, err := uc.GetRolePermissionAssignments("r-1")
	if err != nil {
		t.Fatalf("GetRolePermissionAssignments() error = %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("GetRolePermissionAssignments() len = %d, want %d", len(got), len(want))
	}
}
