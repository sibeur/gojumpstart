package repository

import (
	core_cache "gojumpstart/core/common/cache"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Repository struct {
	User *UserRepository
	Todo *TodoRepository
}

func NewRepository(gormDB *gorm.DB, mongoDB *mongo.Database, cache core_cache.CacheUseCase) *Repository {
	user := NewUserRepository(gormDB)
	todo := NewTodoRepository(mongoDB, cache)
	return &Repository{
		User: user,
		Todo: todo,
	}
}
