package main

import (
	"hexashop/conf"
	"hexashop/internal/adapter/http"
	"hexashop/internal/adapter/http/api"
	"hexashop/internal/adapter/storage/postgres"
	"hexashop/internal/adapter/storage/postgres/repo"
	"hexashop/internal/core/service"
	"log"
)

func main() {
	conf.Setup()

	// Database configuration
	db := postgres.New(conf.DbUrl)
	defer db.Close()

	// Migration
	err := db.MigrateUp()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// User configuration
	userRepo := repo.NewUser(db)
	userSvc := service.NewUser(userRepo)
	userApi := api.NewUser(userSvc)

	// Product configuration
	productRepo := repo.NewProduct(db)
	productSvc := service.NewProduct(productRepo)
	productApi := api.NewProduct(productSvc)

	//Router configuration
	router := http.NewRouter(userApi, productApi)

	log.Println("Server is running on port", conf.HttpPort)
	router.Servce(conf.HttpPort)
}
