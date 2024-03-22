package db

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection() (*mongo.Database, error) {
	uri := os.Getenv("MONGO_DSN")
	if uri == "" {
		return nil, errors.New("MONGO_DSN is required")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	dbName := os.Getenv("MONGO_DBNAME")
	if dbName == "" {
		return nil, errors.New("MONGO_DBNAME is required")
	}

	database := client.Database(dbName)

	var result bson.M
	if err := database.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return nil, err
	}

	return database, nil
}
