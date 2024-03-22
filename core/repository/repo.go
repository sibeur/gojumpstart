package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Repository struct {
	User *UserRepository
	Todo *TodoRepository
}

func NewRepository(gormDB *gorm.DB, mongoDB *mongo.Database) *Repository {
	user := NewUserRepository(gormDB)
	todo := NewTodoRepository(mongoDB)
	return &Repository{
		User: user,
		Todo: todo,
	}
}
