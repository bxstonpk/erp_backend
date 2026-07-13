package entity

type CompanyAccess struct {
	Id        string
	CompanyId string
	BranchId  string
	Default   bool
	OnlyRead  bool
	Active    bool
}
