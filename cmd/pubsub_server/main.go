package main

import (
	"context"
	core_db "gojumpstart/core/db"
	core_repository "gojumpstart/core/repository"
	core_service "gojumpstart/core/service"

	app_pubsub "gojumpstart/apps/pubsub"

	"github.com/joho/godotenv"
)

func main() {
	// This is the entry point of the application
	// dotenv load
	godotenv.Load()

	// database connection
	gormDB, err := core_db.NewMySQLConnection()
	if err != nil {
		panic(err)
	}

	// close db connection
	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	//optional to auto migrate
	// core_db.AutoMigrate(db)

	// load mongodb
	mongoDB, err := core_db.NewMongoDBConnection()
	if err != nil {
		panic(err)
	}
	defer mongoDB.Client().Disconnect(context.Background())

	// load reapository
	repo := core_repository.NewRepository(gormDB, mongoDB)

	// load service
	service := core_service.NewService(repo)

	// start pubsub app
	app_pubsub.NewPubSubApp(service).Run()
}
