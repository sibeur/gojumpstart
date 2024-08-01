package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"gojumpstart/core/entity"
	"log"

	go_cache "github.com/sibeur/go-cache"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	db    *mongo.Database
	cache go_cache.Cache
}

func NewTodoRepository(db *mongo.Database, cache go_cache.Cache) *TodoRepository {
	return &TodoRepository{db: db, cache: cache}
}

func (u *TodoRepository) FindAll() ([]*entity.Todo, error) {
	ctx := context.TODO()
	var todos []*entity.Todo
	cacheDataString, err := u.cache.Get("todos")
	if err == nil && cacheDataString != "" {
		var cacheData []*entity.Todo

		err = json.Unmarshal([]byte(cacheDataString), &cacheData)
		if err != nil {
			return nil, err
		}
		todos = cacheData
		return todos, nil
	}
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

	go func(data []*entity.Todo, cache go_cache.Cache) {
		// Set the data in the cache as json
		cacheDataJSON, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshalling data")
		}

		err = cache.Set("todos", string(cacheDataJSON))
		if err != nil {
			fmt.Println("Failed to set data in cache")
		}
	}(todos, u.cache)

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
