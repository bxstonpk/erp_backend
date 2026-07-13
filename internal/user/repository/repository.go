package repository

import (
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/repository/model"
)

type CommandRepository interface {
	CreateUser(user *entity.User) (error, *entity.User)
	CreateRole(role *entity.Role) (error, *entity.Role)
	CreatePermission(permission *entity.Permission) (error, *entity.Permission)
	AssignUserPermission(userPermission *entity.User) (error, *entity.User)
	AssignUserRole(userRole *entity.User) (error, *entity.User)
	AssignRolePermission(rolePermission *model.RolePermission) (error, *model.RolePermission)
	AssignUserCompanyAccess(userCompanyAccess *model.UserCompanyAccess) (error, *model.UserCompanyAccess)

	UpdateUser(User *entity.User) (error, []*entity.User)
	UpdateRole(Role *entity.Role) (error, []*entity.Role)
	UpdatePermission(Permission *entity.Permission) (error, []*entity.Permission)
	UpdateUserPermission(UserPermission *entity.Permission) (error, []*entity.Permission)
	UpdateUserRole(UserId string, RoleId string) (error, []*entity.Role)
	UpdateRolePermission(RolePermission *entity.Permission) (error, []*entity.Permission)
	UpdateUserCompanyAccess(UserId string, CompanyAccess *entity.CompanyAccess) (error, []*entity.CompanyAccess)

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
	GetAllUserPermissions() ([]*entity.Permission, error)
	GetAllRolePermissions() ([]*entity.Permission, error)
	GetAllUserCompanyAccesses() ([]*entity.CompanyAccess, error)
}

type Repository interface {
	CommandRepository
	QueryRepository
}
