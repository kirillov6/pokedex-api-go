package pokedex

type (
	User struct {
		Id       uint   `json:"id" gorm:"primarykey"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Passowrd string `json:"password"`
	}

	UserPokedex struct {
		Id        uint `gorm:"primarykey"`
		UserId    uint
		PokedexId uint
	}
)
