package postgres

import (
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/repository/model"
)

func (p postgresUserRepository) CreateUser(user *model.User) error {
	return p.db.Create(user).Error
}

func (p postgresUserRepository) CreateRole(role *model.Role) error {
	return p.db.Create(role).Error
}

func (p postgresUserRepository) CreatePermission(permission *model.Permission) error {
	return p.db.Create(permission).Error
}

func (p postgresUserRepository) AssignUserPermission(userPermission *model.UserPermission) error {
	return p.db.Create(userPermission).Error
}

func (p postgresUserRepository) AssignUserRole(userRole *model.UserRole) error {
	return p.db.Create(userRole).Error
}

func (p postgresUserRepository) AssignRolePermission(rolePermission *model.RolePermission) error {
	return p.db.Create(rolePermission).Error
}

func (p postgresUserRepository) AssignUserCompanyAccess(userCompanyAccess *model.UserCompanyAccess) error {
	return p.db.Create(userCompanyAccess).Error
}

func (p postgresUserRepository) UpdateUser(user *model.User) error {
	return p.db.Save(user).Error
}

func (p postgresUserRepository) UpdateRole(role *model.Role) error {
	return p.db.Save(role).Error
}

func (p postgresUserRepository) UpdatePermission(permission *model.Permission) error {
	return p.db.Save(permission).Error
}

func (p postgresUserRepository) UpdateUserPermission(userPermission *model.UserPermission) error {
	return p.db.Save(userPermission).Error
}

func (p postgresUserRepository) UpdateUserRole(userRole *model.UserRole) error {
	return p.db.Save(userRole).Error
}

func (p postgresUserRepository) UpdateRolePermission(rolePermission *model.RolePermission) error {
	return p.db.Save(rolePermission).Error
}

func (p postgresUserRepository) UpdateUserCompanyAccess(userCompanyAccess *model.UserCompanyAccess) error {
	return p.db.Save(userCompanyAccess).Error
}

func (p postgresUserRepository) DeleteUser(userUUID string) error {
	return p.db.Where("user_uuid = ?", userUUID).Delete(&entity.User{}).Error
}

func (p postgresUserRepository) DeleteRole(roleUUID string) error {
	return p.db.Where("role_uuid = ?", roleUUID).Delete(&entity.Role{}).Error
}

func (p postgresUserRepository) DeletePermission(permissionUUID string) error {
	return p.db.Where("permission_uuid = ?", permissionUUID).Delete(&entity.Permission{}).Error
}

func (p postgresUserRepository) DeleteUserPermission(userPermissionUUID string) error {
	return p.db.Where("user_permission_uuid = ?", userPermissionUUID).Delete(&entity.Permission{}).Error
}

func (p postgresUserRepository) DeleteUserRole(userRoleUUID string) error {
	return p.db.Where("user_role_uuid = ?", userRoleUUID).Delete(&entity.Role{}).Error
}

func (p postgresUserRepository) DeleteRolePermission(rolePermissionUUID string) error {
	return p.db.Where("role_permission_uuid = ?", rolePermissionUUID).Delete(&entity.Permission{}).Error
}

func (p postgresUserRepository) DeleteUserCompanyAccess(userCompanyAccessUUID string) error {
	return p.db.Where("user_company_uuid = ?", userCompanyAccessUUID).Delete(&entity.CompanyAccess{}).Error
}
