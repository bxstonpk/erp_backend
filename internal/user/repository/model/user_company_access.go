package model

type UserCompanyAccess struct {
	UserCompanyUUID string `db:"uuid" gorm:"primaryKey" json:"user_company_uuid"`
	UserUUID        string `db:"user_uuid" json:"user_uuid"`
	CompanyUUID     string `db:"company_uuid" json:"company_uuid"`
	BranchUUID      string `db:"branch_uuid" json:"branch_uuid"`
	IsDefault       bool   `db:"is_default" json:"is_default"`
	ReadOnly        bool   `db:"access_only" json:"access_only"`
	Active          bool   `db:"active" json:"active"`
}
