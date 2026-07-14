package model

type UserPermission struct {
	UserPermissionUUID string `db:"uuid" gorm:"primaryKey" json:"user_permission_uuid"`
	UserUUID           string `db:"user_uuid" json:"user_uuid"`
	PermissionUUID     string `db:"permission_uuid" json:"permission_uuid"`
}
