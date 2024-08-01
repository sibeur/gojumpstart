package db_test

import (
	"gojumpstart/core/db"
	"os"
	"testing"
)

func TestNewPostgresConnection(t *testing.T) {
	// Set up test environment
	// ...
	os.Setenv("POSTGRES_DSN", "postgres://postgres:root@localhost:5432/test?sslmode=disable")
	// Call the function being tested
	db, err := db.NewPostgresConnection()
	// Check if the function returned an error
	if err != nil {
		t.Errorf("NewPostgresConnection() returned an error: %v", err)
	}
	// Check if the returned database connection is not nil
	if db == nil {
		t.Error("NewPostgresConnection() returned a nil database connection")
	}
}
