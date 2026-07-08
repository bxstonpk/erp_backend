package model

type UserRole struct {
	UserRoleUUID string `db:"uuid" gorm:"primaryKey" json:"user_role_uuid"`
	UserUUID     string `db:"user_uuid" json:"user_uuid"`
	RoleUUID     string `db:"role_uuid" json:"role_uuid"`
}
