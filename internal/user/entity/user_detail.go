package entity

type User struct {
	UserUUID          string
	CompanyUUID       string
	Username          string
	Email             string
	FullName          string
	Phone             string
	Status            string
	UserRoles         []Role
	UserPermissions   []UserPermission
	RolePermissions   []RolePermission
	UserCompanyAccess []CompanyAccess
}

type Role struct {
	UserRoleUUID string
	UserUUID     string
	RoleUUID     string
}

type UserPermission struct {
	UserPermissionUUID string
	UserUUID           string
	PermissionUUID     string
}

type RolePermission struct {
	RolePermissionUUID string
	RoleUUID           string
	PermissionUUID     string
}

type CompanyAccess struct {
	UserCompanyAccessUUID string
	UserUUID              string
	CompanyUUID           string
}
