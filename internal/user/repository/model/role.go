package model

type Role struct {
	RoleUUID    string `db:"uuid" gorm:"primaryKey" json:"role_uuid"`
	CompanyUUID string `db:"company_uuid" json:"company_uuid"`
	RoleName    string `db:"role_name" gorm:"unique" json:"role_name"`
	RoleDesc    string `db:"role_desc" json:"role_desc"`
}
