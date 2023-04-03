package pokedex

type (
	Creditionals struct {
		Username string `json:"username" binding:"required"`
		Passowrd string `json:"password" binding:"required"`
	}

	User struct {
		Id   uint   `json:"-" gorm:"primarykey"`
		Name string `json:"name" binding:"required"`
		Creditionals
	}

	UserPokedex struct {
		Id        uint `gorm:"primarykey"`
		UserId    uint
		PokedexId uint
	}
)
