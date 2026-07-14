package dto

import "erp/backend/internal/user/entity"

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
}

type UserResponse struct {
	ID              string                  `json:"user_uuid"`
	Username        string                  `json:"username"`
	Email           string                  `json:"email"`
	FullName        string                  `json:"full_name"`
	Phone           string                  `json:"phone"`
	Status          string                  `json:"status"`
	Roles           []RoleResponse          `json:"roles,omitempty"`
	Permissions     []PermissionResponse    `json:"permissions,omitempty"`
	CompanyAccesses []CompanyAccessResponse `json:"company_accesses,omitempty"`
}

func NewUserResponse(u *entity.User) UserResponse {
	roles := make([]RoleResponse, 0, len(u.Roles))
	for i := range u.Roles {
		roles = append(roles, NewRoleResponse(&u.Roles[i]))
	}

	permissions := make([]PermissionResponse, 0, len(u.Permissions))
	for i := range u.Permissions {
		permissions = append(permissions, NewPermissionResponse(&u.Permissions[i]))
	}

	accesses := make([]CompanyAccessResponse, 0, len(u.CompanyAccesses))
	for i := range u.CompanyAccesses {
		accesses = append(accesses, NewCompanyAccessResponse(&u.CompanyAccesses[i]))
	}

	return UserResponse{
		ID:              u.ID,
		Username:        u.Username,
		Email:           u.Email,
		FullName:        u.FullName,
		Phone:           u.Phone,
		Status:          u.Status,
		Roles:           roles,
		Permissions:     permissions,
		CompanyAccesses: accesses,
	}
}

func NewUserResponses(users []*entity.User) []UserResponse {
	res := make([]UserResponse, 0, len(users))
	for _, u := range users {
		res = append(res, NewUserResponse(u))
	}

	return res
}
