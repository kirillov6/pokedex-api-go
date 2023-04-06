package service

import (
	"errors"

	"github.com/kirillov6/pokedex-api-go/pkg/pokedex"
	"github.com/kirillov6/pokedex-api-go/pkg/repository"
)

type PokedexService struct {
	repo repository.Pokedex
}

func NewPokedexService(repo repository.Pokedex) *PokedexService {
	return &PokedexService{repo}
}

func (s *PokedexService) Create(userId uint, pokemon pokedex.Pokemon) (uint, error) {
	return s.repo.Create(userId, pokemon)
}

func (s *PokedexService) GetPokemons(userId uint) ([]pokedex.Pokemon, error) {
	return s.repo.GetPokemons(userId)
}

func (s *PokedexService) GetPokemon(userId, pokemonId uint) (pokedex.Pokemon, error) {
	return s.repo.GetPokemon(userId, pokemonId)
}

func (s *PokedexService) DeletePokemon(userId, pokemonId uint) error {
	return s.repo.DeletePokemon(userId, pokemonId)
}

func (s *PokedexService) UpdatePokemon(userId, pokemonId uint, newData pokedex.Pokemon) error {
	if newData.Name == nil && newData.Type == nil && newData.HP == nil &&
		newData.Attack == nil && newData.Defense == nil && newData.Speed == nil {
		return errors.New("Empty data")
	}
	return s.repo.UpdatePokemon(userId, pokemonId, newData)
}
