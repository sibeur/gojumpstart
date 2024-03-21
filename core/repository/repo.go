package repository

import "gorm.io/gorm"

type Repository struct {
	User *UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	user := NewUserRepository(db)
	return &Repository{
		User: user,
	}
}
