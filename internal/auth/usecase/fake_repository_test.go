package usecase_test

import (
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/repository"
)

var _ repository.QueryRepository = (*fakeQueryRepository)(nil)

// fakeQueryRepository implements repository.QueryRepository with one
// overridable function field per method used by the auth usecase; methods
// the auth usecase doesn't call are left unset and will panic if invoked,
// which flags an incorrect test expectation.
type fakeQueryRepository struct {
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

func (f *fakeQueryRepository) GetUserByUUID(userUUID string) (*entity.User, error) {
	return f.GetUserByUUIDFn(userUUID)
}
func (f *fakeQueryRepository) GetUserByUsername(username string) (*entity.User, error) {
	return f.GetUserByUsernameFn(username)
}
func (f *fakeQueryRepository) GetAllUsers() ([]*entity.User, error) { return f.GetAllUsersFn() }
func (f *fakeQueryRepository) GetAllRoles() ([]*entity.Role, error) { return f.GetAllRolesFn() }
func (f *fakeQueryRepository) GetAllPermissions() ([]*entity.Permission, error) {
	return f.GetAllPermissionsFn()
}
func (f *fakeQueryRepository) GetRoleByUUID(roleUUID string) (*entity.Role, error) {
	return f.GetRoleByUUIDFn(roleUUID)
}
func (f *fakeQueryRepository) GetPermissionByUUID(permissionUUID string) (*entity.Permission, error) {
	return f.GetPermissionByUUIDFn(permissionUUID)
}
func (f *fakeQueryRepository) GetUserRoles(userUUID string) ([]*entity.Role, error) {
	return f.GetUserRolesFn(userUUID)
}
func (f *fakeQueryRepository) GetUserCompanyAccesses(userUUID string) ([]*entity.CompanyAccess, error) {
	return f.GetUserCompanyAccessesFn(userUUID)
}
func (f *fakeQueryRepository) GetPermissionAssignmentsByUser(userUUID string) ([]*entity.PermissionAssignment, error) {
	return f.GetPermissionAssignmentsByUserFn(userUUID)
}
func (f *fakeQueryRepository) GetPermissionAssignmentsByRole(roleUUID string) ([]*entity.PermissionAssignment, error) {
	return f.GetPermissionAssignmentsByRoleFn(roleUUID)
}
