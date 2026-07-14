package dto

import "erp/backend/internal/user/entity"

type CreatePermissionRequest struct {
	CompanyID   string `json:"company_uuid"`
	Code        string `json:"permission_code"`
	Module      string `json:"permission_module"`
	Name        string `json:"permission_name"`
	Description string `json:"permission_desc"`
}

type UpdatePermissionRequest struct {
	CompanyID   string `json:"company_uuid"`
	Code        string `json:"permission_code"`
	Module      string `json:"permission_module"`
	Name        string `json:"permission_name"`
	Description string `json:"permission_desc"`
}

type PermissionResponse struct {
	ID          string `json:"permission_uuid"`
	CompanyID   string `json:"company_uuid"`
	Code        string `json:"permission_code"`
	Module      string `json:"permission_module"`
	Name        string `json:"permission_name"`
	Description string `json:"permission_desc"`
}

func NewPermissionResponse(p *entity.Permission) PermissionResponse {
	return PermissionResponse{
		ID:          p.ID,
		CompanyID:   p.CompanyID,
		Code:        p.Code,
		Module:      p.Module,
		Name:        p.Name,
		Description: p.Description,
	}
}

func NewPermissionResponses(permissions []*entity.Permission) []PermissionResponse {
	res := make([]PermissionResponse, 0, len(permissions))
	for _, p := range permissions {
		res = append(res, NewPermissionResponse(p))
	}

	return res
}
