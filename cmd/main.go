package main

import (
	"log"

	initializer "github.com/kirillov6/pokedex-api-go/pkg"
	"github.com/kirillov6/pokedex-api-go/pkg/handler"
	"github.com/kirillov6/pokedex-api-go/pkg/pokedex"
	"github.com/kirillov6/pokedex-api-go/pkg/repository"
	"github.com/kirillov6/pokedex-api-go/pkg/service"
)

func init() {
	if err := initializer.LoadEnvVariables(); err != nil {
		log.Fatal("Error while read env variables")
	}
}

func main() {
	db, err := initializer.ConnectToDB()
	if err != nil {
		log.Fatalf("Error while connecting to DB: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	server := pokedex.NewServer(handler.InitRoutes())
	if err := server.Run(); err != nil {
		log.Fatalf("Error while running server: %s", err.Error())
	}
}
