package dto_test

import (
	"testing"

	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
)

func TestNewCompanyAccessResponse(t *testing.T) {
	access := &entity.CompanyAccess{CompanyID: "c-1", BranchID: "b-1", IsDefault: true, ReadOnly: true, Active: true}

	got := dto.NewCompanyAccessResponse(access)

	if got.CompanyID != access.CompanyID || got.BranchID != access.BranchID ||
		got.IsDefault != access.IsDefault || got.ReadOnly != access.ReadOnly || got.Active != access.Active {
		t.Fatalf("NewCompanyAccessResponse() = %+v, want fields from %+v", got, access)
	}
}

func TestNewCompanyAccessResponses(t *testing.T) {
	accesses := []*entity.CompanyAccess{
		{CompanyID: "c-1"},
		{CompanyID: "c-2"},
	}

	got := dto.NewCompanyAccessResponses(accesses)

	if len(got) != 2 {
		t.Fatalf("NewCompanyAccessResponses() len = %d, want 2", len(got))
	}
	if got[0].CompanyID != "c-1" || got[1].CompanyID != "c-2" {
		t.Fatalf("NewCompanyAccessResponses() = %+v", got)
	}
}

func TestNewCompanyAccessResponses_empty(t *testing.T) {
	got := dto.NewCompanyAccessResponses(nil)
	if len(got) != 0 {
		t.Fatalf("NewCompanyAccessResponses(nil) len = %d, want 0", len(got))
	}
}
