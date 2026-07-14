package dto

import "erp/backend/internal/user/entity"

type AssignUserRoleRequest struct {
	UserUUID string `json:"user_uuid"`
	RoleUUID string `json:"role_uuid"`
}

type AssignUserPermissionRequest struct {
	UserUUID       string `json:"user_uuid"`
	PermissionUUID string `json:"permission_uuid"`
}

type AssignRolePermissionRequest struct {
	RoleUUID       string `json:"role_uuid"`
	PermissionUUID string `json:"permission_uuid"`
}

type PermissionAssignmentResponse struct {
	AssignmentID string             `json:"assignment_uuid"`
	Source       string             `json:"source"`
	OwnerID      string             `json:"owner_uuid"`
	Permission   PermissionResponse `json:"permission"`
}

func NewPermissionAssignmentResponse(a *entity.PermissionAssignment) PermissionAssignmentResponse {
	return PermissionAssignmentResponse{
		AssignmentID: a.AssignmentID,
		Source:       string(a.Source),
		OwnerID:      a.OwnerID,
		Permission:   NewPermissionResponse(&a.Permission),
	}
}

func NewPermissionAssignmentResponses(assignments []*entity.PermissionAssignment) []PermissionAssignmentResponse {
	res := make([]PermissionAssignmentResponse, 0, len(assignments))
	for _, a := range assignments {
		res = append(res, NewPermissionAssignmentResponse(a))
	}

	return res
}
