package repository

import (
	go_cache "github.com/sibeur/go-cache"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Repository struct {
	User *UserRepository
	Todo *TodoRepository
}

func NewRepository(gormDB *gorm.DB, mongoDB *mongo.Database, cache go_cache.Cache) *Repository {
	user := NewUserRepository(gormDB)
	todo := NewTodoRepository(mongoDB, cache)
	return &Repository{
		User: user,
		Todo: todo,
	}
}
