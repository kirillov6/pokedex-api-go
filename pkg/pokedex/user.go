package pokedex

type (
	Creditionals struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	User struct {
		Id   uint   `json:"-" gorm:"primarykey"`
		Name string `json:"name" binding:"required"`
		Creditionals
	}

	UserPokemon struct {
		Id        uint `gorm:"primarykey"`
		UserId    uint
		PokemonId uint
	}
)
