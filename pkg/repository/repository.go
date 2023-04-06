package repository

import (
	"github.com/kirillov6/pokedex-api-go/pkg/pokedex"
	"gorm.io/gorm"
)

type (
	Authorization interface {
		CreateUser(user pokedex.User) (uint, error)
		GetUser(creds pokedex.Creditionals) (pokedex.User, error)
	}

	Pokedex interface {
		Create(userId uint, pokemon pokedex.Pokemon) (uint, error)
		GetPokemons(userId uint) ([]pokedex.Pokemon, error)
		GetPokemon(userId, pokemonId uint) (pokedex.Pokemon, error)
		DeletePokemon(userId, pokemonId uint) error
		UpdatePokemon(userId, pokemonId uint, newData pokedex.Pokemon) error
	}

	Repository struct {
		Authorization
		Pokedex
	}
)

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Pokedex:       NewPokedexRepository(db),
	}
}
