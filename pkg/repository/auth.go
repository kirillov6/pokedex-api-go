package repository

import (
	"github.com/kirillov6/pokedex-api-go/pkg/pokedex"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateUser(user pokedex.User) (uint, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (r *AuthRepository) GetUser(creds pokedex.Creditionals) (pokedex.User, error) {
	var user pokedex.User
	res := r.db.Where(&pokedex.User{Creditionals: creds}).First(&user)
	return user, res.Error
}
