package postgres

import (
	"gorm.io/gorm"
)

type postgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *postgresUserRepository {
	return &postgresUserRepository{
		db: db,
	}
}
