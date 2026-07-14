package model

type Permission struct {
	PermissionUUID   string `db:"uuid" gorm:"primaryKey" json:"permission_uuid"`
	PermissionCode   string `db:"permission_code" gorm:"unique" json:"permission_code"`
	CompanyUUID      string `db:"company_uuid" json:"company_uuid"`
	PermissionName   string `db:"permission_name" gorm:"unique" json:"permission_name"`
	PermissionModule string `db:"permission_module" json:"permission_module"`
	PermissionDesc   string `db:"permission_desc" json:"permission_desc"`
}
