package handler_test

import (
	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
	"erp/backend/internal/user/usecase"
)

// fakeUsecase implements usecase.Usecase with one overridable function field
// per method, so each handler test wires up only the calls it expects.
type fakeUsecase struct {
	CreateUserFn        func(req dto.CreateUserRequest) (*entity.User, error)
	GetUserFn           func(userUUID string) (*entity.User, error)
	GetUserByUsernameFn func(username string) (*entity.User, error)
	ListUsersFn         func() ([]*entity.User, error)
	UpdateUserFn        func(userUUID string, req dto.UpdateUserRequest) (*entity.User, error)
	DeleteUserFn        func(userUUID string) error

	CreateRoleFn func(req dto.CreateRoleRequest) (*entity.Role, error)
	GetRoleFn    func(roleUUID string) (*entity.Role, error)
	ListRolesFn  func() ([]*entity.Role, error)
	UpdateRoleFn func(roleUUID string, req dto.UpdateRoleRequest) (*entity.Role, error)
	DeleteRoleFn func(roleUUID string) error

	CreatePermissionFn func(req dto.CreatePermissionRequest) (*entity.Permission, error)
	GetPermissionFn    func(permissionUUID string) (*entity.Permission, error)
	ListPermissionsFn  func() ([]*entity.Permission, error)
	UpdatePermissionFn func(permissionUUID string, req dto.UpdatePermissionRequest) (*entity.Permission, error)
	DeletePermissionFn func(permissionUUID string) error

	AssignUserRoleFn       func(req dto.AssignUserRoleRequest) error
	AssignUserPermissionFn func(req dto.AssignUserPermissionRequest) error
	AssignRolePermissionFn func(req dto.AssignRolePermissionRequest) error

	UpdateUserRoleFn       func(userRoleUUID string, req dto.AssignUserRoleRequest) error
	UpdateUserPermissionFn func(userPermissionUUID string, req dto.AssignUserPermissionRequest) error
	UpdateRolePermissionFn func(rolePermissionUUID string, req dto.AssignRolePermissionRequest) error

	DeleteUserRoleFn       func(userRoleUUID string) error
	DeleteUserPermissionFn func(userPermissionUUID string) error
	DeleteRolePermissionFn func(rolePermissionUUID string) error

	GetUserRolesFn                 func(userUUID string) ([]*entity.Role, error)
	GetUserPermissionAssignmentsFn func(userUUID string) ([]*entity.PermissionAssignment, error)
	GetRolePermissionAssignmentsFn func(roleUUID string) ([]*entity.PermissionAssignment, error)

	AssignUserCompanyAccessFn func(userUUID string, req dto.CompanyAccessRequest) error
	UpdateUserCompanyAccessFn func(userCompanyUUID string, req dto.UpdateCompanyAccessRequest) error
	DeleteUserCompanyAccessFn func(userCompanyAccessUUID string) error
	GetUserCompanyAccessesFn  func(userUUID string) ([]*entity.CompanyAccess, error)
}

var _ usecase.Usecase = (*fakeUsecase)(nil)

func (f *fakeUsecase) CreateUser(req dto.CreateUserRequest) (*entity.User, error) {
	return f.CreateUserFn(req)
}
func (f *fakeUsecase) GetUser(userUUID string) (*entity.User, error) { return f.GetUserFn(userUUID) }
func (f *fakeUsecase) GetUserByUsername(username string) (*entity.User, error) {
	return f.GetUserByUsernameFn(username)
}
func (f *fakeUsecase) ListUsers() ([]*entity.User, error) { return f.ListUsersFn() }
func (f *fakeUsecase) UpdateUser(userUUID string, req dto.UpdateUserRequest) (*entity.User, error) {
	return f.UpdateUserFn(userUUID, req)
}
func (f *fakeUsecase) DeleteUser(userUUID string) error { return f.DeleteUserFn(userUUID) }

func (f *fakeUsecase) CreateRole(req dto.CreateRoleRequest) (*entity.Role, error) {
	return f.CreateRoleFn(req)
}
func (f *fakeUsecase) GetRole(roleUUID string) (*entity.Role, error) { return f.GetRoleFn(roleUUID) }
func (f *fakeUsecase) ListRoles() ([]*entity.Role, error)            { return f.ListRolesFn() }
func (f *fakeUsecase) UpdateRole(roleUUID string, req dto.UpdateRoleRequest) (*entity.Role, error) {
	return f.UpdateRoleFn(roleUUID, req)
}
func (f *fakeUsecase) DeleteRole(roleUUID string) error { return f.DeleteRoleFn(roleUUID) }

func (f *fakeUsecase) CreatePermission(req dto.CreatePermissionRequest) (*entity.Permission, error) {
	return f.CreatePermissionFn(req)
}
func (f *fakeUsecase) GetPermission(permissionUUID string) (*entity.Permission, error) {
	return f.GetPermissionFn(permissionUUID)
}
func (f *fakeUsecase) ListPermissions() ([]*entity.Permission, error) { return f.ListPermissionsFn() }
func (f *fakeUsecase) UpdatePermission(permissionUUID string, req dto.UpdatePermissionRequest) (*entity.Permission, error) {
	return f.UpdatePermissionFn(permissionUUID, req)
}
func (f *fakeUsecase) DeletePermission(permissionUUID string) error {
	return f.DeletePermissionFn(permissionUUID)
}

func (f *fakeUsecase) AssignUserRole(req dto.AssignUserRoleRequest) error {
	return f.AssignUserRoleFn(req)
}
func (f *fakeUsecase) AssignUserPermission(req dto.AssignUserPermissionRequest) error {
	return f.AssignUserPermissionFn(req)
}
func (f *fakeUsecase) AssignRolePermission(req dto.AssignRolePermissionRequest) error {
	return f.AssignRolePermissionFn(req)
}
func (f *fakeUsecase) UpdateUserRole(userRoleUUID string, req dto.AssignUserRoleRequest) error {
	return f.UpdateUserRoleFn(userRoleUUID, req)
}
func (f *fakeUsecase) UpdateUserPermission(userPermissionUUID string, req dto.AssignUserPermissionRequest) error {
	return f.UpdateUserPermissionFn(userPermissionUUID, req)
}
func (f *fakeUsecase) UpdateRolePermission(rolePermissionUUID string, req dto.AssignRolePermissionRequest) error {
	return f.UpdateRolePermissionFn(rolePermissionUUID, req)
}
func (f *fakeUsecase) DeleteUserRole(userRoleUUID string) error {
	return f.DeleteUserRoleFn(userRoleUUID)
}
func (f *fakeUsecase) DeleteUserPermission(userPermissionUUID string) error {
	return f.DeleteUserPermissionFn(userPermissionUUID)
}
func (f *fakeUsecase) DeleteRolePermission(rolePermissionUUID string) error {
	return f.DeleteRolePermissionFn(rolePermissionUUID)
}
func (f *fakeUsecase) GetUserRoles(userUUID string) ([]*entity.Role, error) {
	return f.GetUserRolesFn(userUUID)
}
func (f *fakeUsecase) GetUserPermissionAssignments(userUUID string) ([]*entity.PermissionAssignment, error) {
	return f.GetUserPermissionAssignmentsFn(userUUID)
}
func (f *fakeUsecase) GetRolePermissionAssignments(roleUUID string) ([]*entity.PermissionAssignment, error) {
	return f.GetRolePermissionAssignmentsFn(roleUUID)
}

func (f *fakeUsecase) AssignUserCompanyAccess(userUUID string, req dto.CompanyAccessRequest) error {
	return f.AssignUserCompanyAccessFn(userUUID, req)
}
func (f *fakeUsecase) UpdateUserCompanyAccess(userCompanyUUID string, req dto.UpdateCompanyAccessRequest) error {
	return f.UpdateUserCompanyAccessFn(userCompanyUUID, req)
}
func (f *fakeUsecase) DeleteUserCompanyAccess(userCompanyAccessUUID string) error {
	return f.DeleteUserCompanyAccessFn(userCompanyAccessUUID)
}
func (f *fakeUsecase) GetUserCompanyAccesses(userUUID string) ([]*entity.CompanyAccess, error) {
	return f.GetUserCompanyAccessesFn(userUUID)
}
