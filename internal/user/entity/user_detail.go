package entity

type User struct {
	Id              string
	Username        string
	Email           string
	FullName        string
	Phone           string
	Status          string
	Roles           []Role
	Permissions     []Permission
	RolePermissions []RolePermission
	CompanyAccess   []CompanyAccess
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
