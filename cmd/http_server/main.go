package main

import (
	app_http "gojumpstart/apps/http"
	core_db "gojumpstart/core/db"
	core_repository "gojumpstart/core/repository"
	core_service "gojumpstart/core/service"

	go_cache "github.com/sibeur/go-cache"

	"github.com/joho/godotenv"
)

func main() {
	// This is the entry point of the application
	// dotenv load
	godotenv.Load()

	// database connection
	gormDB, err := core_db.NewPostgresConnection()
	if err != nil {
		panic(err)
	}

	// close db connection
	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	// optional to auto migrate
	core_db.AutoMigrate(gormDB)

	cache := go_cache.NewCache()

	// load reapository
	repo := core_repository.NewRepository(gormDB, cache)

	// load service
	service := core_service.NewService(repo)

	// start http app
	app_http.NewFiberApp(service).Run()

}
