package db_test

import (
	"context"
	"gojumpstart/core/db"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestNewMongoDbConnection(t *testing.T) {
	// Set up environment variables
	os.Setenv("MONGO_DSN", "mongodb://localhost:27017")
	os.Setenv("MONGO_DBNAME", "testdb")

	// Clean up environment variables after the test
	defer func() {
		os.Unsetenv("MONGO_DSN")
		os.Unsetenv("MONGO_DBNAME")
	}()

	// Call the function under test
	database, err := db.NewMongoDBConnection()
	defer database.Client().Disconnect(context.Background())

	// Check if there was an error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the returned database is not nil
	if database == nil {
		t.Error("Expected non-nil database, got nil")
	}

	// Check if the ping command was successful
	var result bson.M
	if err := database.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		t.Errorf("Unexpected error while pinging database: %v", err)
	}
}
