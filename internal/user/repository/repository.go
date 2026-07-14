package repository

import (
	"erp/backend/internal/user/entity"
)

type CommandRepository interface {
	CreateUser(user *entity.User) error
	CreateRole(role *entity.Role) error
	CreatePermission(permission *entity.Permission) error
	AssignUserPermission(userUUID, permissionUUID string) error
	AssignUserRole(userUUID, roleUUID string) error
	AssignRolePermission(roleUUID, permissionUUID string) error
	AssignUserCompanyAccess(userUUID string, access *entity.CompanyAccess) error

	UpdateUser(user *entity.User) error
	UpdateRole(role *entity.Role) error
	UpdatePermission(permission *entity.Permission) error
	UpdateUserPermission(userPermissionUUID, userUUID, permissionUUID string) error
	UpdateUserRole(userRoleUUID, userUUID, roleUUID string) error
	UpdateRolePermission(rolePermissionUUID, roleUUID, permissionUUID string) error
	UpdateUserCompanyAccess(userCompanyUUID, userUUID string, access *entity.CompanyAccess) error

	DeleteUser(userUUID string) error
	DeleteRole(roleUUID string) error
	DeletePermission(permissionUUID string) error
	DeleteUserPermission(userPermissionUUID string) error
	DeleteUserRole(userRoleUUID string) error
	DeleteRolePermission(rolePermissionUUID string) error
	DeleteUserCompanyAccess(userCompanyAccessUUID string) error
}

type QueryRepository interface {
	GetUserByUUID(userUUID string) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)

	GetAllUsers() ([]*entity.User, error)
	GetAllRoles() ([]*entity.Role, error)
	GetAllPermissions() ([]*entity.Permission, error)

	GetRoleByUUID(roleUUID string) (*entity.Role, error)
	GetPermissionByUUID(permissionUUID string) (*entity.Permission, error)

	GetUserRoles(userUUID string) ([]*entity.Role, error)
	GetUserCompanyAccesses(userUUID string) ([]*entity.CompanyAccess, error)

	GetPermissionAssignmentsByUser(userUUID string) ([]*entity.PermissionAssignment, error)
	GetPermissionAssignmentsByRole(roleUUID string) ([]*entity.PermissionAssignment, error)
}

type Repository interface {
	CommandRepository
	QueryRepository
}
