package model

type User struct {
	UserUUID       string `db:"uuid" gorm:"primaryKey" json:"user_uuid"`
	CompanyUUID    string `db:"company_uuid" json:"company_uuid"`
	Username       string `db:"username" gorm:"unique" json:"username"`
	Email          string `db:"email" gorm:"unique" json:"email"`
	HashedPassword string `db:"hashed_password" json:"password"`
	FullName       string `db:"full_name" json:"full_name"`
	Phone          string `db:"phone" json:"phone"`
	Status         string `db:"status" json:"status"`
}
