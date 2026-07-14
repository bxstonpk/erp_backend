package entity

type User struct {
	ID       string
	Username string
	Email    string
	Password string
	FullName string
	Phone    string
	Status   string

	Roles           []Role
	Permissions     []Permission
	CompanyAccesses []CompanyAccess
}
