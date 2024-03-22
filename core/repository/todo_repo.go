package repository

import (
	"context"
	"gojumpstart/core/entity"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	db *mongo.Database
}

func NewTodoRepository(db *mongo.Database) *TodoRepository {
	return &TodoRepository{db: db}
}

func (u *TodoRepository) FindAll() ([]*entity.Todo, error) {
	ctx := context.TODO()
	var todos []*entity.Todo
	cur, err := u.db.Collection(entity.Todo{}.GetCollName()).Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Error finding todos: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var todo entity.Todo
		err := cur.Decode(&todo)
		if err != nil {
			log.Printf("Error decoding todo: %v", err)
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (u *TodoRepository) Create(todo *entity.Todo) error {
	todo.ID = primitive.NewObjectID()
	_, err := u.db.Collection(entity.Todo{}.GetCollName()).InsertOne(context.TODO(), todo)
	if err != nil {
		return err
	}
	return nil
}
