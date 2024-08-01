package repository

import (
	core_cache "gojumpstart/core/common/cache"

	"gorm.io/gorm"
)

type Repository struct {
	User *UserRepository
}

func NewRepository(gormDB *gorm.DB, cache core_cache.CacheUseCase) *Repository {
	user := NewUserRepository(gormDB)
	return &Repository{
		User: user,
	}
}
