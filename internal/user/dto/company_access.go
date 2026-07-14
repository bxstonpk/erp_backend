package dto

import "erp/backend/internal/user/entity"

type CompanyAccessRequest struct {
	CompanyID string `json:"company_uuid"`
	BranchID  string `json:"branch_uuid"`
	IsDefault bool   `json:"is_default"`
	ReadOnly  bool   `json:"access_only"`
	Active    bool   `json:"active"`
}

// UpdateCompanyAccessRequest carries UserUUID because UpdateUserCompanyAccess
// re-targets the assignment and the entity itself has no owning-user field.
type UpdateCompanyAccessRequest struct {
	UserUUID  string `json:"user_uuid"`
	CompanyID string `json:"company_uuid"`
	BranchID  string `json:"branch_uuid"`
	IsDefault bool   `json:"is_default"`
	ReadOnly  bool   `json:"access_only"`
	Active    bool   `json:"active"`
}

type CompanyAccessResponse struct {
	CompanyID string `json:"company_uuid"`
	BranchID  string `json:"branch_uuid"`
	IsDefault bool   `json:"is_default"`
	ReadOnly  bool   `json:"access_only"`
	Active    bool   `json:"active"`
}

func NewCompanyAccessResponse(a *entity.CompanyAccess) CompanyAccessResponse {
	return CompanyAccessResponse{
		CompanyID: a.CompanyID,
		BranchID:  a.BranchID,
		IsDefault: a.IsDefault,
		ReadOnly:  a.ReadOnly,
		Active:    a.Active,
	}
}

func NewCompanyAccessResponses(accesses []*entity.CompanyAccess) []CompanyAccessResponse {
	res := make([]CompanyAccessResponse, 0, len(accesses))
	for _, a := range accesses {
		res = append(res, NewCompanyAccessResponse(a))
	}

	return res
}
