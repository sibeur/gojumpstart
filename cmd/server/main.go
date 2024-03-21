package main

import (
	app_http "gojumpstart/apps/http"
	core_db "gojumpstart/core/db"
	core_repository "gojumpstart/core/repository"
	core_service "gojumpstart/core/service"

	"github.com/joho/godotenv"
)

func main() {
	// This is the entry point of the application
	// dotenv load
	godotenv.Load()

	// database connection
	db, err := core_db.NewMySQLConnection()
	if err != nil {
		panic(err)
	}

	// close db connection
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	//optional to auto migrate
	core_db.AutoMigrate(db)

	// load reapository
	repo := core_repository.NewRepository(db)

	// load service
	service := core_service.NewService(repo)

	// start http app
	app_http.NewFiberApp(service).Run()

}
