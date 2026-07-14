package repository

import (
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/repository/model"
)

// UserToEntity assembles the business User aggregate from the persistence
// User row plus its related rows, which the repository loads separately
// since the model layer has no knowledge of the business shape.
func UserToEntity(m *model.User, roles []model.Role, permissions []model.Permission, accesses []model.UserCompanyAccess) *entity.User {
	if m == nil {
		return nil
	}

	return &entity.User{
		ID:              m.UserUUID,
		Username:        m.Username,
		Email:           m.Email,
		FullName:        m.FullName,
		Phone:           m.Phone,
		Status:          m.Status,
		Roles:           RolesToEntities(roles),
		Permissions:     PermissionsToEntities(permissions),
		CompanyAccesses: CompanyAccessesToEntities(accesses),
	}
}

func RoleToEntity(m model.Role) entity.Role {
	return entity.Role{
		ID:          m.RoleUUID,
		CompanyID:   m.CompanyUUID,
		Name:        m.RoleName,
		Description: m.RoleDesc,
	}
}

func RolesToEntities(models []model.Role) []entity.Role {
	roles := make([]entity.Role, 0, len(models))
	for _, m := range models {
		roles = append(roles, RoleToEntity(m))
	}

	return roles
}

func PermissionToEntity(m model.Permission) entity.Permission {
	return entity.Permission{
		ID:          m.PermissionUUID,
		CompanyID:   m.CompanyUUID,
		Code:        m.PermissionCode,
		Module:      m.PermissionModule,
		Name:        m.PermissionName,
		Description: m.PermissionDesc,
	}
}

func PermissionsToEntities(models []model.Permission) []entity.Permission {
	permissions := make([]entity.Permission, 0, len(models))
	for _, m := range models {
		permissions = append(permissions, PermissionToEntity(m))
	}

	return permissions
}

func CompanyAccessToEntity(m model.UserCompanyAccess) entity.CompanyAccess {
	return entity.CompanyAccess{
		CompanyID: m.CompanyUUID,
		BranchID:  m.BranchUUID,
		IsDefault: m.IsDefault,
		ReadOnly:  m.ReadOnly,
		Active:    m.Active,
	}
}

func CompanyAccessesToEntities(models []model.UserCompanyAccess) []entity.CompanyAccess {
	accesses := make([]entity.CompanyAccess, 0, len(models))
	for _, m := range models {
		accesses = append(accesses, CompanyAccessToEntity(m))
	}

	return accesses
}

// UserToModel maps only the User row itself. Roles, Permissions and
// CompanyAccesses are persisted separately through the Assign* commands,
// since they live in their own junction/lookup tables.
func UserToModel(e *entity.User) *model.User {
	if e == nil {
		return nil
	}

	return &model.User{
		UserUUID:       e.ID,
		Username:       e.Username,
		Email:          e.Email,
		HashedPassword: e.Password,
		FullName:       e.FullName,
		Phone:          e.Phone,
		Status:         e.Status,
	}
}

func RoleToModel(e entity.Role) model.Role {
	return model.Role{
		RoleUUID:    e.ID,
		CompanyUUID: e.CompanyID,
		RoleName:    e.Name,
		RoleDesc:    e.Description,
	}
}

func PermissionToModel(e entity.Permission) model.Permission {
	return model.Permission{
		PermissionUUID:   e.ID,
		CompanyUUID:      e.CompanyID,
		PermissionCode:   e.Code,
		PermissionModule: e.Module,
		PermissionName:   e.Name,
		PermissionDesc:   e.Description,
	}
}

// CompanyAccessToModel takes userUUID separately since entity.CompanyAccess
// is nested under entity.User and doesn't carry a back-reference to it.
func CompanyAccessToModel(userUUID string, e entity.CompanyAccess) model.UserCompanyAccess {
	return model.UserCompanyAccess{
		UserUUID:    userUUID,
		CompanyUUID: e.CompanyID,
		BranchUUID:  e.BranchID,
		IsDefault:   e.IsDefault,
		ReadOnly:    e.ReadOnly,
		Active:      e.Active,
	}
}

func UserPermissionToAssignment(m model.UserPermission, permission model.Permission) entity.PermissionAssignment {
	return entity.PermissionAssignment{
		AssignmentID: m.UserPermissionUUID,
		Source:       entity.PermissionSourceUser,
		OwnerID:      m.UserUUID,
		Permission:   PermissionToEntity(permission),
	}
}

func RolePermissionToAssignment(m model.RolePermission, permission model.Permission) entity.PermissionAssignment {
	return entity.PermissionAssignment{
		AssignmentID: m.RolePermissionUUID,
		Source:       entity.PermissionSourceRole,
		OwnerID:      m.RoleUUID,
		Permission:   PermissionToEntity(permission),
	}
}
