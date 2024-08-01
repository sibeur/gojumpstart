package repository

import (
	go_cache "github.com/sibeur/go-cache"

	"gorm.io/gorm"
)

type Repository struct {
	User *UserRepository
}

func NewRepository(gormDB *gorm.DB, cache go_cache.Cache) *Repository {
	user := NewUserRepository(gormDB)
	return &Repository{
		User: user,
	}
}
