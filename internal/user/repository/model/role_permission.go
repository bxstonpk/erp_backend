package model

type RolePermission struct {
	RolePermissionUUID string `db:"uuid" gorm:"primaryKey" json:"role_permission_uuid"`
	RoleUUID           string `db:"role_uuid" json:"role_uuid"`
	PermissionUUID     string `db:"permission_uuid" json:"permission_uuid"`
}
