package service

import (
	"github.com/kirillov6/pokedex-api-go/pkg/pokedex"
	"github.com/kirillov6/pokedex-api-go/pkg/repository"
)

type (
	Authorization interface {
		CreateUser(user pokedex.User) (uint, error)
		GenerateToken(creds pokedex.Creditionals) (string, error)
		ParseToken(token string) (uint, error)
	}

	Pokedex interface {
		Create(userId uint, pokemon pokedex.Pokemon) (uint, error)
		GetPokemons(userId uint) ([]pokedex.Pokemon, error)
		GetPokemon(userId, pokemonId uint) (pokedex.Pokemon, error)
		DeletePokemon(userId, pokemonId uint) error
		UpdatePokemon(userId, pokemonId uint, newData pokedex.Pokemon) error
	}

	Service struct {
		Authorization
		Pokedex
	}
)

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Pokedex:       NewPokedexService(repo.Pokedex),
	}
}
