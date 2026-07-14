package usecase_test

import (
	"errors"
	"testing"

	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/usecase"
)

func TestPermissionUsecase_CreatePermission(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var captured *entity.Permission
		repo := &fakeRepository{
			CreatePermissionFn: func(permission *entity.Permission) error {
				captured = permission
				return nil
			},
		}
		uc := usecase.New(repo)

		req := dto.CreatePermissionRequest{CompanyID: "c-1", Code: "user.read", Module: "user", Name: "Read User", Description: "view users"}
		got, err := uc.CreatePermission(req)
		if err != nil {
			t.Fatalf("CreatePermission() error = %v", err)
		}
		if got.ID == "" {
			t.Fatal("CreatePermission() did not generate an ID")
		}
		if captured != got {
			t.Fatal("CreatePermission() did not persist the returned permission")
		}
		if got.CompanyID != req.CompanyID || got.Code != req.Code || got.Module != req.Module ||
			got.Name != req.Name || got.Description != req.Description {
			t.Fatalf("CreatePermission() mismatched fields: %+v", got)
		}
	})

	t.Run("repository error", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			CreatePermissionFn: func(permission *entity.Permission) error { return wantErr },
		}
		uc := usecase.New(repo)

		if _, err := uc.CreatePermission(dto.CreatePermissionRequest{}); !errors.Is(err, wantErr) {
			t.Fatalf("CreatePermission() error = %v, want %v", err, wantErr)
		}
	})
}

func TestPermissionUsecase_GetPermission(t *testing.T) {
	want := &entity.Permission{ID: "p-1", Code: "user.read"}
	repo := &fakeRepository{
		GetPermissionByUUIDFn: func(permissionUUID string) (*entity.Permission, error) {
			if permissionUUID != "p-1" {
				t.Fatalf("GetPermissionByUUID() permissionUUID = %q, want %q", permissionUUID, "p-1")
			}
			return want, nil
		},
	}
	uc := usecase.New(repo)

	got, err := uc.GetPermission("p-1")
	if err != nil {
		t.Fatalf("GetPermission() error = %v", err)
	}
	if got != want {
		t.Fatalf("GetPermission() = %v, want %v", got, want)
	}
}

func TestPermissionUsecase_ListPermissions(t *testing.T) {
	want := []*entity.Permission{{ID: "p-1"}, {ID: "p-2"}}
	repo := &fakeRepository{
		GetAllPermissionsFn: func() ([]*entity.Permission, error) { return want, nil },
	}
	uc := usecase.New(repo)

	got, err := uc.ListPermissions()
	if err != nil {
		t.Fatalf("ListPermissions() error = %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("ListPermissions() len = %d, want %d", len(got), len(want))
	}
}

func TestPermissionUsecase_UpdatePermission(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var saved *entity.Permission
		repo := &fakeRepository{
			UpdatePermissionFn: func(permission *entity.Permission) error {
				saved = permission
				return nil
			},
		}
		uc := usecase.New(repo)

		req := dto.UpdatePermissionRequest{CompanyID: "c-1", Code: "user.write", Module: "user", Name: "Write User", Description: "edit users"}
		got, err := uc.UpdatePermission("p-1", req)
		if err != nil {
			t.Fatalf("UpdatePermission() error = %v", err)
		}
		if got.ID != "p-1" || got.CompanyID != req.CompanyID || got.Code != req.Code ||
			got.Module != req.Module || got.Name != req.Name || got.Description != req.Description {
			t.Fatalf("UpdatePermission() mismatched fields: %+v", got)
		}
		if saved != got {
			t.Fatal("UpdatePermission() did not persist the updated permission")
		}
	})

	t.Run("repository error", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			UpdatePermissionFn: func(permission *entity.Permission) error { return wantErr },
		}
		uc := usecase.New(repo)

		if _, err := uc.UpdatePermission("p-1", dto.UpdatePermissionRequest{}); !errors.Is(err, wantErr) {
			t.Fatalf("UpdatePermission() error = %v, want %v", err, wantErr)
		}
	})
}

func TestPermissionUsecase_DeletePermission(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var got string
		repo := &fakeRepository{
			DeletePermissionFn: func(permissionUUID string) error {
				got = permissionUUID
				return nil
			},
		}
		uc := usecase.New(repo)

		if err := uc.DeletePermission("p-1"); err != nil {
			t.Fatalf("DeletePermission() error = %v", err)
		}
		if got != "p-1" {
			t.Fatalf("DeletePermission() permissionUUID = %q, want %q", got, "p-1")
		}
	})

	t.Run("repository error", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			DeletePermissionFn: func(permissionUUID string) error { return wantErr },
		}
		uc := usecase.New(repo)

		if err := uc.DeletePermission("p-1"); !errors.Is(err, wantErr) {
			t.Fatalf("DeletePermission() error = %v, want %v", err, wantErr)
		}
	})
}
