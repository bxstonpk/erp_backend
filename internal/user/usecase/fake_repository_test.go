package usecase_test

import (
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/repository"
)

var _ repository.Repository = (*fakeRepository)(nil)

// fakeRepository implements repository.Repository with one overridable
// function field per method, so each test wires up only the calls it
// expects and can assert on the arguments it receives.
type fakeRepository struct {
	CreateUserFn              func(user *entity.User) error
	CreateRoleFn              func(role *entity.Role) error
	CreatePermissionFn        func(permission *entity.Permission) error
	AssignUserPermissionFn    func(userUUID, permissionUUID string) error
	AssignUserRoleFn          func(userUUID, roleUUID string) error
	AssignRolePermissionFn    func(roleUUID, permissionUUID string) error
	AssignUserCompanyAccessFn func(userUUID string, access *entity.CompanyAccess) error

	UpdateUserFn              func(user *entity.User) error
	UpdateRoleFn              func(role *entity.Role) error
	UpdatePermissionFn        func(permission *entity.Permission) error
	UpdateUserPermissionFn    func(userPermissionUUID, userUUID, permissionUUID string) error
	UpdateUserRoleFn          func(userRoleUUID, userUUID, roleUUID string) error
	UpdateRolePermissionFn    func(rolePermissionUUID, roleUUID, permissionUUID string) error
	UpdateUserCompanyAccessFn func(userCompanyUUID, userUUID string, access *entity.CompanyAccess) error

	DeleteUserFn              func(userUUID string) error
	DeleteRoleFn              func(roleUUID string) error
	DeletePermissionFn        func(permissionUUID string) error
	DeleteUserPermissionFn    func(userPermissionUUID string) error
	DeleteUserRoleFn          func(userRoleUUID string) error
	DeleteRolePermissionFn    func(rolePermissionUUID string) error
	DeleteUserCompanyAccessFn func(userCompanyAccessUUID string) error

	GetUserByUUIDFn                  func(userUUID string) (*entity.User, error)
	GetUserByUsernameFn              func(username string) (*entity.User, error)
	GetAllUsersFn                    func() ([]*entity.User, error)
	GetAllRolesFn                    func() ([]*entity.Role, error)
	GetAllPermissionsFn              func() ([]*entity.Permission, error)
	GetRoleByUUIDFn                  func(roleUUID string) (*entity.Role, error)
	GetPermissionByUUIDFn            func(permissionUUID string) (*entity.Permission, error)
	GetUserRolesFn                   func(userUUID string) ([]*entity.Role, error)
	GetUserCompanyAccessesFn         func(userUUID string) ([]*entity.CompanyAccess, error)
	GetPermissionAssignmentsByUserFn func(userUUID string) ([]*entity.PermissionAssignment, error)
	GetPermissionAssignmentsByRoleFn func(roleUUID string) ([]*entity.PermissionAssignment, error)
}

func (f *fakeRepository) CreateUser(user *entity.User) error { return f.CreateUserFn(user) }
func (f *fakeRepository) CreateRole(role *entity.Role) error { return f.CreateRoleFn(role) }
func (f *fakeRepository) CreatePermission(permission *entity.Permission) error {
	return f.CreatePermissionFn(permission)
}
func (f *fakeRepository) AssignUserPermission(userUUID, permissionUUID string) error {
	return f.AssignUserPermissionFn(userUUID, permissionUUID)
}
func (f *fakeRepository) AssignUserRole(userUUID, roleUUID string) error {
	return f.AssignUserRoleFn(userUUID, roleUUID)
}
func (f *fakeRepository) AssignRolePermission(roleUUID, permissionUUID string) error {
	return f.AssignRolePermissionFn(roleUUID, permissionUUID)
}
func (f *fakeRepository) AssignUserCompanyAccess(userUUID string, access *entity.CompanyAccess) error {
	return f.AssignUserCompanyAccessFn(userUUID, access)
}

func (f *fakeRepository) UpdateUser(user *entity.User) error { return f.UpdateUserFn(user) }
func (f *fakeRepository) UpdateRole(role *entity.Role) error { return f.UpdateRoleFn(role) }
func (f *fakeRepository) UpdatePermission(permission *entity.Permission) error {
	return f.UpdatePermissionFn(permission)
}
func (f *fakeRepository) UpdateUserPermission(userPermissionUUID, userUUID, permissionUUID string) error {
	return f.UpdateUserPermissionFn(userPermissionUUID, userUUID, permissionUUID)
}
func (f *fakeRepository) UpdateUserRole(userRoleUUID, userUUID, roleUUID string) error {
	return f.UpdateUserRoleFn(userRoleUUID, userUUID, roleUUID)
}
func (f *fakeRepository) UpdateRolePermission(rolePermissionUUID, roleUUID, permissionUUID string) error {
	return f.UpdateRolePermissionFn(rolePermissionUUID, roleUUID, permissionUUID)
}
func (f *fakeRepository) UpdateUserCompanyAccess(userCompanyUUID, userUUID string, access *entity.CompanyAccess) error {
	return f.UpdateUserCompanyAccessFn(userCompanyUUID, userUUID, access)
}

func (f *fakeRepository) DeleteUser(userUUID string) error { return f.DeleteUserFn(userUUID) }
func (f *fakeRepository) DeleteRole(roleUUID string) error { return f.DeleteRoleFn(roleUUID) }
func (f *fakeRepository) DeletePermission(permissionUUID string) error {
	return f.DeletePermissionFn(permissionUUID)
}
func (f *fakeRepository) DeleteUserPermission(userPermissionUUID string) error {
	return f.DeleteUserPermissionFn(userPermissionUUID)
}
func (f *fakeRepository) DeleteUserRole(userRoleUUID string) error {
	return f.DeleteUserRoleFn(userRoleUUID)
}
func (f *fakeRepository) DeleteRolePermission(rolePermissionUUID string) error {
	return f.DeleteRolePermissionFn(rolePermissionUUID)
}
func (f *fakeRepository) DeleteUserCompanyAccess(userCompanyAccessUUID string) error {
	return f.DeleteUserCompanyAccessFn(userCompanyAccessUUID)
}

func (f *fakeRepository) GetUserByUUID(userUUID string) (*entity.User, error) {
	return f.GetUserByUUIDFn(userUUID)
}
func (f *fakeRepository) GetUserByUsername(username string) (*entity.User, error) {
	return f.GetUserByUsernameFn(username)
}
func (f *fakeRepository) GetAllUsers() ([]*entity.User, error) { return f.GetAllUsersFn() }
func (f *fakeRepository) GetAllRoles() ([]*entity.Role, error) { return f.GetAllRolesFn() }
func (f *fakeRepository) GetAllPermissions() ([]*entity.Permission, error) {
	return f.GetAllPermissionsFn()
}
func (f *fakeRepository) GetRoleByUUID(roleUUID string) (*entity.Role, error) {
	return f.GetRoleByUUIDFn(roleUUID)
}
func (f *fakeRepository) GetPermissionByUUID(permissionUUID string) (*entity.Permission, error) {
	return f.GetPermissionByUUIDFn(permissionUUID)
}
func (f *fakeRepository) GetUserRoles(userUUID string) ([]*entity.Role, error) {
	return f.GetUserRolesFn(userUUID)
}
func (f *fakeRepository) GetUserCompanyAccesses(userUUID string) ([]*entity.CompanyAccess, error) {
	return f.GetUserCompanyAccessesFn(userUUID)
}
func (f *fakeRepository) GetPermissionAssignmentsByUser(userUUID string) ([]*entity.PermissionAssignment, error) {
	return f.GetPermissionAssignmentsByUserFn(userUUID)
}
func (f *fakeRepository) GetPermissionAssignmentsByRole(roleUUID string) ([]*entity.PermissionAssignment, error) {
	return f.GetPermissionAssignmentsByRoleFn(roleUUID)
}
