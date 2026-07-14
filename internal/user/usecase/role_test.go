package usecase_test

import (
	"errors"
	"testing"

	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/usecase"
)

func TestRoleUsecase_CreateRole(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var captured *entity.Role
		repo := &fakeRepository{
			CreateRoleFn: func(role *entity.Role) error {
				captured = role
				return nil
			},
		}
		uc := usecase.New(repo)

		req := dto.CreateRoleRequest{CompanyID: "c-1", Name: "admin", Description: "administrator"}
		got, err := uc.CreateRole(req)
		if err != nil {
			t.Fatalf("CreateRole() error = %v", err)
		}
		if got.ID == "" {
			t.Fatal("CreateRole() did not generate an ID")
		}
		if captured != got {
			t.Fatal("CreateRole() did not persist the returned role")
		}
		if got.CompanyID != req.CompanyID || got.Name != req.Name || got.Description != req.Description {
			t.Fatalf("CreateRole() mismatched fields: %+v", got)
		}
	})

	t.Run("repository error", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			CreateRoleFn: func(role *entity.Role) error { return wantErr },
		}
		uc := usecase.New(repo)

		if _, err := uc.CreateRole(dto.CreateRoleRequest{}); !errors.Is(err, wantErr) {
			t.Fatalf("CreateRole() error = %v, want %v", err, wantErr)
		}
	})
}

func TestRoleUsecase_GetRole(t *testing.T) {
	want := &entity.Role{ID: "r-1", Name: "admin"}
	repo := &fakeRepository{
		GetRoleByUUIDFn: func(roleUUID string) (*entity.Role, error) {
			if roleUUID != "r-1" {
				t.Fatalf("GetRoleByUUID() roleUUID = %q, want %q", roleUUID, "r-1")
			}
			return want, nil
		},
	}
	uc := usecase.New(repo)

	got, err := uc.GetRole("r-1")
	if err != nil {
		t.Fatalf("GetRole() error = %v", err)
	}
	if got != want {
		t.Fatalf("GetRole() = %v, want %v", got, want)
	}
}

func TestRoleUsecase_ListRoles(t *testing.T) {
	want := []*entity.Role{{ID: "r-1"}, {ID: "r-2"}}
	repo := &fakeRepository{
		GetAllRolesFn: func() ([]*entity.Role, error) { return want, nil },
	}
	uc := usecase.New(repo)

	got, err := uc.ListRoles()
	if err != nil {
		t.Fatalf("ListRoles() error = %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("ListRoles() len = %d, want %d", len(got), len(want))
	}
}

func TestRoleUsecase_UpdateRole(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var saved *entity.Role
		repo := &fakeRepository{
			UpdateRoleFn: func(role *entity.Role) error {
				saved = role
				return nil
			},
		}
		uc := usecase.New(repo)

		req := dto.UpdateRoleRequest{CompanyID: "c-1", Name: "admin", Description: "administrator"}
		got, err := uc.UpdateRole("r-1", req)
		if err != nil {
			t.Fatalf("UpdateRole() error = %v", err)
		}
		if got.ID != "r-1" || got.CompanyID != req.CompanyID || got.Name != req.Name || got.Description != req.Description {
			t.Fatalf("UpdateRole() mismatched fields: %+v", got)
		}
		if saved != got {
			t.Fatal("UpdateRole() did not persist the updated role")
		}
	})

	t.Run("repository error", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			UpdateRoleFn: func(role *entity.Role) error { return wantErr },
		}
		uc := usecase.New(repo)

		if _, err := uc.UpdateRole("r-1", dto.UpdateRoleRequest{}); !errors.Is(err, wantErr) {
			t.Fatalf("UpdateRole() error = %v, want %v", err, wantErr)
		}
	})
}

func TestRoleUsecase_DeleteRole(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var got string
		repo := &fakeRepository{
			DeleteRoleFn: func(roleUUID string) error {
				got = roleUUID
				return nil
			},
		}
		uc := usecase.New(repo)

		if err := uc.DeleteRole("r-1"); err != nil {
			t.Fatalf("DeleteRole() error = %v", err)
		}
		if got != "r-1" {
			t.Fatalf("DeleteRole() roleUUID = %q, want %q", got, "r-1")
		}
	})

	t.Run("repository error", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			DeleteRoleFn: func(roleUUID string) error { return wantErr },
		}
		uc := usecase.New(repo)

		if err := uc.DeleteRole("r-1"); !errors.Is(err, wantErr) {
			t.Fatalf("DeleteRole() error = %v, want %v", err, wantErr)
		}
	})
}
