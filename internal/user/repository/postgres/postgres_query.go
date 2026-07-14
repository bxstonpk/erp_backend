package postgres

import (
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/repository"
	"erp/backend/internal/user/repository/model"
)

func (r *postgresUserRepository) loadUserAggregate(userModel *model.User) (*entity.User, error) {
	var userRoles []model.UserRole
	if err := r.db.Where("user_uuid = ?", userModel.UserUUID).Find(&userRoles).Error; err != nil {
		return nil, err
	}

	roleUUIDs := make([]string, 0, len(userRoles))
	for _, userRole := range userRoles {
		roleUUIDs = append(roleUUIDs, userRole.RoleUUID)
	}

	var roles []model.Role
	if len(roleUUIDs) > 0 {
		if err := r.db.Where("role_uuid IN ?", roleUUIDs).Find(&roles).Error; err != nil {
			return nil, err
		}
	}

	var userPermissions []model.UserPermission
	if err := r.db.Where("user_uuid = ?", userModel.UserUUID).Find(&userPermissions).Error; err != nil {
		return nil, err
	}

	permissionUUIDs := make([]string, 0, len(userPermissions))
	for _, userPermission := range userPermissions {
		permissionUUIDs = append(permissionUUIDs, userPermission.PermissionUUID)
	}

	var permissions []model.Permission
	if len(permissionUUIDs) > 0 {
		if err := r.db.Where("permission_uuid IN ?", permissionUUIDs).Find(&permissions).Error; err != nil {
			return nil, err
		}
	}

	var companyAccesses []model.UserCompanyAccess
	if err := r.db.Where("user_uuid = ?", userModel.UserUUID).Find(&companyAccesses).Error; err != nil {
		return nil, err
	}

	return repository.UserToEntity(userModel, roles, permissions, companyAccesses), nil
}

func (r *postgresUserRepository) GetUserByUUID(userUUID string) (*entity.User, error) {
	userModel := &model.User{}
	if err := r.db.Where("user_uuid = ?", userUUID).First(userModel).Error; err != nil {
		return nil, err
	}

	return r.loadUserAggregate(userModel)
}

func (r *postgresUserRepository) GetUserByUsername(username string) (*entity.User, error) {
	userModel := &model.User{}
	if err := r.db.Where("username = ?", username).First(userModel).Error; err != nil {
		return nil, err
	}

	return r.loadUserAggregate(userModel)
}

func (r *postgresUserRepository) GetAllUsers() ([]*entity.User, error) {
	var userModels []model.User
	if err := r.db.Find(&userModels).Error; err != nil {
		return nil, err
	}

	users := make([]*entity.User, 0, len(userModels))
	for i := range userModels {
		user, err := r.loadUserAggregate(&userModels[i])
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *postgresUserRepository) GetAllRoles() ([]*entity.Role, error) {
	var roleModels []model.Role
	if err := r.db.Find(&roleModels).Error; err != nil {
		return nil, err
	}

	roles := make([]*entity.Role, 0, len(roleModels))
	for _, roleModel := range roleModels {
		role := repository.RoleToEntity(roleModel)
		roles = append(roles, &role)
	}

	return roles, nil
}

func (r *postgresUserRepository) GetAllPermissions() ([]*entity.Permission, error) {
	var permissionModels []model.Permission
	if err := r.db.Find(&permissionModels).Error; err != nil {
		return nil, err
	}

	permissions := make([]*entity.Permission, 0, len(permissionModels))
	for _, permissionModel := range permissionModels {
		permission := repository.PermissionToEntity(permissionModel)
		permissions = append(permissions, &permission)
	}

	return permissions, nil
}

func (r *postgresUserRepository) GetRoleByUUID(roleUUID string) (*entity.Role, error) {
	roleModel := model.Role{}
	if err := r.db.Where("role_uuid = ?", roleUUID).First(&roleModel).Error; err != nil {
		return nil, err
	}

	role := repository.RoleToEntity(roleModel)
	return &role, nil
}

func (r *postgresUserRepository) GetPermissionByUUID(permissionUUID string) (*entity.Permission, error) {
	permissionModel := model.Permission{}
	if err := r.db.Where("permission_uuid = ?", permissionUUID).First(&permissionModel).Error; err != nil {
		return nil, err
	}

	permission := repository.PermissionToEntity(permissionModel)
	return &permission, nil
}

func (r *postgresUserRepository) GetUserRoles(userUUID string) ([]*entity.Role, error) {
	var userRoles []model.UserRole
	if err := r.db.Where("user_uuid = ?", userUUID).Find(&userRoles).Error; err != nil {
		return nil, err
	}

	roleUUIDs := make([]string, 0, len(userRoles))
	for _, userRole := range userRoles {
		roleUUIDs = append(roleUUIDs, userRole.RoleUUID)
	}

	var roleModels []model.Role
	if len(roleUUIDs) > 0 {
		if err := r.db.Where("role_uuid IN ?", roleUUIDs).Find(&roleModels).Error; err != nil {
			return nil, err
		}
	}

	roles := make([]*entity.Role, 0, len(roleModels))
	for _, roleModel := range roleModels {
		role := repository.RoleToEntity(roleModel)
		roles = append(roles, &role)
	}

	return roles, nil
}

func (r *postgresUserRepository) GetUserCompanyAccesses(userUUID string) ([]*entity.CompanyAccess, error) {
	var companyAccesses []model.UserCompanyAccess
	if err := r.db.Where("user_uuid = ?", userUUID).Find(&companyAccesses).Error; err != nil {
		return nil, err
	}

	accesses := make([]*entity.CompanyAccess, 0, len(companyAccesses))
	for _, companyAccessModel := range companyAccesses {
		access := repository.CompanyAccessToEntity(companyAccessModel)
		accesses = append(accesses, &access)
	}

	return accesses, nil
}

func (r *postgresUserRepository) permissionsByUUID(permissionUUIDs []string) (map[string]model.Permission, error) {
	permissions := make(map[string]model.Permission, len(permissionUUIDs))
	if len(permissionUUIDs) == 0 {
		return permissions, nil
	}

	var permissionModels []model.Permission
	if err := r.db.Where("permission_uuid IN ?", permissionUUIDs).Find(&permissionModels).Error; err != nil {
		return nil, err
	}

	for _, permissionModel := range permissionModels {
		permissions[permissionModel.PermissionUUID] = permissionModel
	}

	return permissions, nil
}

func (r *postgresUserRepository) GetPermissionAssignmentsByUser(userUUID string) ([]*entity.PermissionAssignment, error) {
	var userPermissions []model.UserPermission
	if err := r.db.Where("user_uuid = ?", userUUID).Find(&userPermissions).Error; err != nil {
		return nil, err
	}

	permissionUUIDs := make([]string, 0, len(userPermissions))
	for _, userPermission := range userPermissions {
		permissionUUIDs = append(permissionUUIDs, userPermission.PermissionUUID)
	}

	permissions, err := r.permissionsByUUID(permissionUUIDs)
	if err != nil {
		return nil, err
	}

	assignments := make([]*entity.PermissionAssignment, 0, len(userPermissions))
	for _, userPermission := range userPermissions {
		assignment := repository.UserPermissionToAssignment(userPermission, permissions[userPermission.PermissionUUID])
		assignments = append(assignments, &assignment)
	}

	return assignments, nil
}

func (r *postgresUserRepository) GetPermissionAssignmentsByRole(roleUUID string) ([]*entity.PermissionAssignment, error) {
	var rolePermissions []model.RolePermission
	if err := r.db.Where("role_uuid = ?", roleUUID).Find(&rolePermissions).Error; err != nil {
		return nil, err
	}

	permissionUUIDs := make([]string, 0, len(rolePermissions))
	for _, rolePermission := range rolePermissions {
		permissionUUIDs = append(permissionUUIDs, rolePermission.PermissionUUID)
	}

	permissions, err := r.permissionsByUUID(permissionUUIDs)
	if err != nil {
		return nil, err
	}

	assignments := make([]*entity.PermissionAssignment, 0, len(rolePermissions))
	for _, rolePermission := range rolePermissions {
		assignment := repository.RolePermissionToAssignment(rolePermission, permissions[rolePermission.PermissionUUID])
		assignments = append(assignments, &assignment)
	}

	return assignments, nil
}
