package repository

import (
	"github.com/kirillov6/pokedex-api-go/pkg/pokedex"
	"gorm.io/gorm"
)

type PokedexRepository struct {
	db *gorm.DB
}

func NewPokedexRepository(db *gorm.DB) *PokedexRepository {
	return &PokedexRepository{db}
}

func (r *PokedexRepository) Create(userId uint, pokemon pokedex.Pokemon) (uint, error) {
	r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&pokemon).Error; err != nil {
			return err
		}

		if err := tx.Create(&pokedex.UserPokemon{
			UserId:    userId,
			PokemonId: pokemon.Id,
		}).Error; err != nil {
			return err
		}

		return nil
	})

	return pokemon.Id, nil
}

func (r *PokedexRepository) GetPokemons(userId uint) ([]pokedex.Pokemon, error) {
	var pokemons []pokedex.Pokemon
	res := r.db.InnerJoins("JOIN user_pokemons on user_pokemons.pokemon_id=pokemons.id").
		Where("user_pokemons.user_id = ?", userId).
		Find(&pokemons)
	return pokemons, res.Error
}

func (r *PokedexRepository) GetPokemon(userId, pokemonId uint) (pokedex.Pokemon, error) {
	var pokemon pokedex.Pokemon
	res := r.db.InnerJoins("JOIN user_pokemons on user_pokemons.pokemon_id=pokemons.id").
		Where("user_pokemons.user_id = ? AND user_pokemons.pokemon_id = ?", userId, pokemonId).
		First(&pokemon)
	return pokemon, res.Error
}

func (r *PokedexRepository) DeletePokemon(userId, pokemonId uint) error {
	pokemon, err := r.GetPokemon(userId, pokemonId)
	if err != nil {
		return err
	}

	res := r.db.Delete(&pokemon)
	return res.Error
}

func (r *PokedexRepository) UpdatePokemon(userId, pokemonId uint, newData pokedex.Pokemon) error {
	pokemon, err := r.GetPokemon(userId, pokemonId)
	if err != nil {
		return err
	}

	if newData.Name != nil {
		pokemon.Name = newData.Name
	}

	if newData.Type != nil {
		pokemon.Type = newData.Type
	}

	if newData.HP != nil {
		pokemon.HP = newData.HP
	}

	if newData.Attack != nil {
		pokemon.Attack = newData.Attack
	}

	if newData.Defense != nil {
		pokemon.Defense = newData.Defense
	}

	if newData.Speed != nil {
		pokemon.Speed = newData.Speed
	}

	res := r.db.Save(&pokemon)
	return res.Error
}
