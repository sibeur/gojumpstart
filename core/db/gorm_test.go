package db_test

import (
	"gojumpstart/core/db"
	"os"
	"testing"
)

func TestNewMySQLConnection(t *testing.T) {
	// Set up test environment
	// ...
	os.Setenv("MYSQL_DSN", "root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")

	// Call the function being tested
	db, err := db.NewMySQLConnection()

	// Check if the function returned an error
	if err != nil {
		t.Errorf("NewMySQLConnection() returned an error: %v", err)
	}

	// Check if the returned database connection is not nil
	if db == nil {
		t.Error("NewMySQLConnection() returned a nil database connection")
	}

}
