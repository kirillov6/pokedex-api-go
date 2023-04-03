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

func (s *AuthRepository) CreateUser(user pokedex.User) (uint, error) {
	res := s.db.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}

	return user.Id, nil
}

func (s *AuthRepository) GetUser(creds pokedex.Creditionals) (pokedex.User, error) {
	var user pokedex.User
	res := s.db.Where(&pokedex.User{Creditionals: creds}).First(&user)
	return user, res.Error
}
