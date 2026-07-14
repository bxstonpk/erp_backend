package entity

type CompanyAccess struct {
	CompanyID string
	BranchID  string
	IsDefault bool
	ReadOnly  bool
	Active    bool
}
