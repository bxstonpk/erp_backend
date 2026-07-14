package dto

import "erp/backend/internal/user/entity"

type CreateRoleRequest struct {
	CompanyID   string `json:"company_uuid"`
	Name        string `json:"role_name"`
	Description string `json:"role_desc"`
}

type UpdateRoleRequest struct {
	CompanyID   string `json:"company_uuid"`
	Name        string `json:"role_name"`
	Description string `json:"role_desc"`
}

type RoleResponse struct {
	ID          string `json:"role_uuid"`
	CompanyID   string `json:"company_uuid"`
	Name        string `json:"role_name"`
	Description string `json:"role_desc"`
}

func NewRoleResponse(r *entity.Role) RoleResponse {
	return RoleResponse{
		ID:          r.ID,
		CompanyID:   r.CompanyID,
		Name:        r.Name,
		Description: r.Description,
	}
}

func NewRoleResponses(roles []*entity.Role) []RoleResponse {
	res := make([]RoleResponse, 0, len(roles))
	for _, r := range roles {
		res = append(res, NewRoleResponse(r))
	}

	return res
}
