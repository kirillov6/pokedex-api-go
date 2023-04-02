package main

import (
	"log"

	initializer "github.com/kirillov6/pokedex-api-go/pkg"
	"github.com/kirillov6/pokedex-api-go/pkg/pokedex"
)

func init() {
	if err := initializer.LoadEnvVariables(); err != nil {
		log.Fatal("Error while read env variables")
	}
}

func main() {
	DB, err := initializer.ConnectToDB()
	if err != nil {
		log.Fatal("Error while connecting to DB")
	}

	DB.AutoMigrate(&pokedex.User{})
	DB.AutoMigrate(&pokedex.Pokedex{})
	DB.AutoMigrate(&pokedex.UserPokedex{})
	DB.AutoMigrate(&pokedex.Pokemon{})
	DB.AutoMigrate(&pokedex.PokedexPokemon{})
}
