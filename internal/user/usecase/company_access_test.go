package usecase_test

import (
	"errors"
	"testing"

	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/usecase"
)

func TestCompanyAccessUsecase_AssignUserCompanyAccess(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var gotUser string
		var gotAccess *entity.CompanyAccess
		repo := &fakeRepository{
			AssignUserCompanyAccessFn: func(userUUID string, access *entity.CompanyAccess) error {
				gotUser, gotAccess = userUUID, access
				return nil
			},
		}
		uc := usecase.New(repo)

		req := dto.CompanyAccessRequest{CompanyID: "c-1", BranchID: "b-1", IsDefault: true, ReadOnly: false, Active: true}
		if err := uc.AssignUserCompanyAccess("u-1", req); err != nil {
			t.Fatalf("AssignUserCompanyAccess() error = %v", err)
		}
		if gotUser != "u-1" {
			t.Fatalf("AssignUserCompanyAccess() userUUID = %q, want %q", gotUser, "u-1")
		}
		if gotAccess.CompanyID != req.CompanyID || gotAccess.BranchID != req.BranchID ||
			gotAccess.IsDefault != req.IsDefault || gotAccess.ReadOnly != req.ReadOnly || gotAccess.Active != req.Active {
			t.Fatalf("AssignUserCompanyAccess() mismatched fields: %+v", gotAccess)
		}
	})

	t.Run("repository error", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			AssignUserCompanyAccessFn: func(userUUID string, access *entity.CompanyAccess) error { return wantErr },
		}
		uc := usecase.New(repo)

		if err := uc.AssignUserCompanyAccess("u-1", dto.CompanyAccessRequest{}); !errors.Is(err, wantErr) {
			t.Fatalf("AssignUserCompanyAccess() error = %v, want %v", err, wantErr)
		}
	})
}

func TestCompanyAccessUsecase_UpdateUserCompanyAccess(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var gotUserCompany, gotUser string
		var gotAccess *entity.CompanyAccess
		repo := &fakeRepository{
			UpdateUserCompanyAccessFn: func(userCompanyUUID, userUUID string, access *entity.CompanyAccess) error {
				gotUserCompany, gotUser, gotAccess = userCompanyUUID, userUUID, access
				return nil
			},
		}
		uc := usecase.New(repo)

		req := dto.UpdateCompanyAccessRequest{
			UserUUID: "u-1", CompanyID: "c-1", BranchID: "b-1", IsDefault: true, ReadOnly: true, Active: false,
		}
		if err := uc.UpdateUserCompanyAccess("uc-1", req); err != nil {
			t.Fatalf("UpdateUserCompanyAccess() error = %v", err)
		}
		if gotUserCompany != "uc-1" || gotUser != "u-1" {
			t.Fatalf("UpdateUserCompanyAccess() forwarded (%q, %q), want (%q, %q)", gotUserCompany, gotUser, "uc-1", "u-1")
		}
		if gotAccess.CompanyID != req.CompanyID || gotAccess.BranchID != req.BranchID ||
			gotAccess.IsDefault != req.IsDefault || gotAccess.ReadOnly != req.ReadOnly || gotAccess.Active != req.Active {
			t.Fatalf("UpdateUserCompanyAccess() mismatched fields: %+v", gotAccess)
		}
	})

	t.Run("repository error", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			UpdateUserCompanyAccessFn: func(userCompanyUUID, userUUID string, access *entity.CompanyAccess) error {
				return wantErr
			},
		}
		uc := usecase.New(repo)

		if err := uc.UpdateUserCompanyAccess("uc-1", dto.UpdateCompanyAccessRequest{}); !errors.Is(err, wantErr) {
			t.Fatalf("UpdateUserCompanyAccess() error = %v, want %v", err, wantErr)
		}
	})
}

func TestCompanyAccessUsecase_DeleteUserCompanyAccess(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var got string
		repo := &fakeRepository{
			DeleteUserCompanyAccessFn: func(userCompanyAccessUUID string) error {
				got = userCompanyAccessUUID
				return nil
			},
		}
		uc := usecase.New(repo)

		if err := uc.DeleteUserCompanyAccess("uc-1"); err != nil {
			t.Fatalf("DeleteUserCompanyAccess() error = %v", err)
		}
		if got != "uc-1" {
			t.Fatalf("DeleteUserCompanyAccess() userCompanyAccessUUID = %q, want %q", got, "uc-1")
		}
	})

	t.Run("repository error", func(t *testing.T) {
		wantErr := errors.New("db down")
		repo := &fakeRepository{
			DeleteUserCompanyAccessFn: func(userCompanyAccessUUID string) error { return wantErr },
		}
		uc := usecase.New(repo)

		if err := uc.DeleteUserCompanyAccess("uc-1"); !errors.Is(err, wantErr) {
			t.Fatalf("DeleteUserCompanyAccess() error = %v, want %v", err, wantErr)
		}
	})
}

func TestCompanyAccessUsecase_GetUserCompanyAccesses(t *testing.T) {
	want := []*entity.CompanyAccess{{CompanyID: "c-1"}, {CompanyID: "c-2"}}
	repo := &fakeRepository{
		GetUserCompanyAccessesFn: func(userUUID string) ([]*entity.CompanyAccess, error) {
			if userUUID != "u-1" {
				t.Fatalf("GetUserCompanyAccesses() userUUID = %q, want %q", userUUID, "u-1")
			}
			return want, nil
		},
	}
	uc := usecase.New(repo)

	got, err := uc.GetUserCompanyAccesses("u-1")
	if err != nil {
		t.Fatalf("GetUserCompanyAccesses() error = %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("GetUserCompanyAccesses() len = %d, want %d", len(got), len(want))
	}
}
