package pokedex

type (
	Pokedex struct {
		Id          uint   `json:"-" gorm:"primarykey"`
		Description string `json:"description"`
	}

	Pokemon struct {
		Id      uint   `json:"-" gorm:"primarykey"`
		Name    string `json:"name"`
		Type    string `json:"type"`
		HP      uint   `json:"hp"`
		Attack  uint   `json:"attack"`
		Defense uint   `json:"defense"`
		Speed   uint   `json:"speed"`
	}

	PokedexPokemon struct {
		Id        uint `gorm:"primarykey"`
		PokedexId uint
		PokemonId uint
	}
)
