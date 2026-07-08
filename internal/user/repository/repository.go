package repository

import (
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/repository/model"
)

type CommandRepository interface {
	CreateUser(user *model.User) error
	CreateRole(role *model.Role) error
	CreatePermission(permission *model.Permission) error
	AssignUserPermission(userPermission *model.UserPermission) error
	AssignUserRole(userRole *model.UserRole) error
	AssignRolePermission(rolePermission *model.RolePermission) error
	AssignUserCompanyAccess(userCompanyAccess *model.UserCompanyAccess) error

	UpdateUser(user *model.User) error
	UpdateRole(role *model.Role) error
	UpdatePermission(permission *model.Permission) error
	UpdateUserPermission(userPermission *model.UserPermission) error
	UpdateUserRole(userRole *model.UserRole) error
	UpdateRolePermission(rolePermission *model.RolePermission) error
	UpdateUserCompanyAccess(userCompanyAccess *model.UserCompanyAccess) error

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
	GetAllUserRoles() ([]*entity.Role, error)
	GetAllUserPermissions() ([]*entity.UserPermission, error)
	GetAllRolePermissions() ([]*entity.RolePermission, error)
	GetAllUserCompanyAccesses() ([]*entity.CompanyAccess, error)
}

type Repository interface {
	CommandRepository
	QueryRepository
}
