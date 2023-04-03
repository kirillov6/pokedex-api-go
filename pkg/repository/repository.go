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
	}

	Repository struct {
		Authorization
		Pokedex
	}
)

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
