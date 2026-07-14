package postgres

import (
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/repository"
	"erp/backend/internal/user/repository/model"
)

func (r *postgresUserRepository) CreateUser(user *entity.User) error {
	return r.db.Create(repository.UserToModel(user)).Error
}

func (r *postgresUserRepository) CreateRole(role *entity.Role) error {
	roleModel := repository.RoleToModel(*role)
	return r.db.Create(&roleModel).Error
}

func (r *postgresUserRepository) CreatePermission(permission *entity.Permission) error {
	permissionModel := repository.PermissionToModel(*permission)
	return r.db.Create(&permissionModel).Error
}

func (r *postgresUserRepository) AssignUserPermission(userUUID, permissionUUID string) error {
	return r.db.Create(&model.UserPermission{
		UserUUID:       userUUID,
		PermissionUUID: permissionUUID,
	}).Error
}

func (r *postgresUserRepository) AssignUserRole(userUUID, roleUUID string) error {
	return r.db.Create(&model.UserRole{
		UserUUID: userUUID,
		RoleUUID: roleUUID,
	}).Error
}

func (r *postgresUserRepository) AssignRolePermission(roleUUID, permissionUUID string) error {
	return r.db.Create(&model.RolePermission{
		RoleUUID:       roleUUID,
		PermissionUUID: permissionUUID,
	}).Error
}

func (r *postgresUserRepository) AssignUserCompanyAccess(userUUID string, access *entity.CompanyAccess) error {
	accessModel := repository.CompanyAccessToModel(userUUID, *access)
	return r.db.Create(&accessModel).Error
}

func (r *postgresUserRepository) UpdateUser(user *entity.User) error {
	return r.db.Save(repository.UserToModel(user)).Error
}

func (r *postgresUserRepository) UpdateRole(role *entity.Role) error {
	roleModel := repository.RoleToModel(*role)
	return r.db.Save(&roleModel).Error
}

func (r *postgresUserRepository) UpdatePermission(permission *entity.Permission) error {
	permissionModel := repository.PermissionToModel(*permission)
	return r.db.Save(&permissionModel).Error
}

func (r *postgresUserRepository) UpdateUserPermission(userPermissionUUID, userUUID, permissionUUID string) error {
	return r.db.Save(&model.UserPermission{
		UserPermissionUUID: userPermissionUUID,
		UserUUID:           userUUID,
		PermissionUUID:     permissionUUID,
	}).Error
}

func (r *postgresUserRepository) UpdateUserRole(userRoleUUID, userUUID, roleUUID string) error {
	return r.db.Save(&model.UserRole{
		UserRoleUUID: userRoleUUID,
		UserUUID:     userUUID,
		RoleUUID:     roleUUID,
	}).Error
}

func (r *postgresUserRepository) UpdateRolePermission(rolePermissionUUID, roleUUID, permissionUUID string) error {
	return r.db.Save(&model.RolePermission{
		RolePermissionUUID: rolePermissionUUID,
		RoleUUID:           roleUUID,
		PermissionUUID:     permissionUUID,
	}).Error
}

func (r *postgresUserRepository) UpdateUserCompanyAccess(userCompanyUUID, userUUID string, access *entity.CompanyAccess) error {
	accessModel := repository.CompanyAccessToModel(userUUID, *access)
	accessModel.UserCompanyUUID = userCompanyUUID
	return r.db.Save(&accessModel).Error
}

func (r *postgresUserRepository) DeleteUser(userUUID string) error {
	return r.db.Where("user_uuid = ?", userUUID).Delete(&model.User{}).Error
}

func (r *postgresUserRepository) DeleteRole(roleUUID string) error {
	return r.db.Where("role_uuid = ?", roleUUID).Delete(&model.Role{}).Error
}

func (r *postgresUserRepository) DeletePermission(permissionUUID string) error {
	return r.db.Where("permission_uuid = ?", permissionUUID).Delete(&model.Permission{}).Error
}

func (r *postgresUserRepository) DeleteUserPermission(userPermissionUUID string) error {
	return r.db.Where("user_permission_uuid = ?", userPermissionUUID).Delete(&model.UserPermission{}).Error
}

func (r *postgresUserRepository) DeleteUserRole(userRoleUUID string) error {
	return r.db.Where("user_role_uuid = ?", userRoleUUID).Delete(&model.UserRole{}).Error
}

func (r *postgresUserRepository) DeleteRolePermission(rolePermissionUUID string) error {
	return r.db.Where("role_permission_uuid = ?", rolePermissionUUID).Delete(&model.RolePermission{}).Error
}

func (r *postgresUserRepository) DeleteUserCompanyAccess(userCompanyAccessUUID string) error {
	return r.db.Where("user_company_uuid = ?", userCompanyAccessUUID).Delete(&model.UserCompanyAccess{}).Error
}
