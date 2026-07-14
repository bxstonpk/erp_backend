package usecase

import (
	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
)

type AssignmentUsecase interface {
	AssignUserRole(req dto.AssignUserRoleRequest) error
	AssignUserPermission(req dto.AssignUserPermissionRequest) error
	AssignRolePermission(req dto.AssignRolePermissionRequest) error

	UpdateUserRole(userRoleUUID string, req dto.AssignUserRoleRequest) error
	UpdateUserPermission(userPermissionUUID string, req dto.AssignUserPermissionRequest) error
	UpdateRolePermission(rolePermissionUUID string, req dto.AssignRolePermissionRequest) error

	DeleteUserRole(userRoleUUID string) error
	DeleteUserPermission(userPermissionUUID string) error
	DeleteRolePermission(rolePermissionUUID string) error

	GetUserRoles(userUUID string) ([]*entity.Role, error)
	GetUserPermissionAssignments(userUUID string) ([]*entity.PermissionAssignment, error)
	GetRolePermissionAssignments(roleUUID string) ([]*entity.PermissionAssignment, error)
}

func (u *userUsecase) AssignUserRole(req dto.AssignUserRoleRequest) error {
	return u.repo.AssignUserRole(req.UserUUID, req.RoleUUID)
}

func (u *userUsecase) AssignUserPermission(req dto.AssignUserPermissionRequest) error {
	return u.repo.AssignUserPermission(req.UserUUID, req.PermissionUUID)
}

func (u *userUsecase) AssignRolePermission(req dto.AssignRolePermissionRequest) error {
	return u.repo.AssignRolePermission(req.RoleUUID, req.PermissionUUID)
}

func (u *userUsecase) UpdateUserRole(userRoleUUID string, req dto.AssignUserRoleRequest) error {
	return u.repo.UpdateUserRole(userRoleUUID, req.UserUUID, req.RoleUUID)
}

func (u *userUsecase) UpdateUserPermission(userPermissionUUID string, req dto.AssignUserPermissionRequest) error {
	return u.repo.UpdateUserPermission(userPermissionUUID, req.UserUUID, req.PermissionUUID)
}

func (u *userUsecase) UpdateRolePermission(rolePermissionUUID string, req dto.AssignRolePermissionRequest) error {
	return u.repo.UpdateRolePermission(rolePermissionUUID, req.RoleUUID, req.PermissionUUID)
}

func (u *userUsecase) DeleteUserRole(userRoleUUID string) error {
	return u.repo.DeleteUserRole(userRoleUUID)
}

func (u *userUsecase) DeleteUserPermission(userPermissionUUID string) error {
	return u.repo.DeleteUserPermission(userPermissionUUID)
}

func (u *userUsecase) DeleteRolePermission(rolePermissionUUID string) error {
	return u.repo.DeleteRolePermission(rolePermissionUUID)
}

func (u *userUsecase) GetUserRoles(userUUID string) ([]*entity.Role, error) {
	return u.repo.GetUserRoles(userUUID)
}

func (u *userUsecase) GetUserPermissionAssignments(userUUID string) ([]*entity.PermissionAssignment, error) {
	return u.repo.GetPermissionAssignmentsByUser(userUUID)
}

func (u *userUsecase) GetRolePermissionAssignments(roleUUID string) ([]*entity.PermissionAssignment, error) {
	return u.repo.GetPermissionAssignmentsByRole(roleUUID)
}
