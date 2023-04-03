package service

import (
	"github.com/kirillov6/pokedex-api-go/pkg/pokedex"
	"github.com/kirillov6/pokedex-api-go/pkg/repository"
)

type (
	Authorization interface {
		CreateUser(user pokedex.User) (uint, error)
		GenerateToken(creds pokedex.Creditionals) (string, error)
		ParseToken(token string) (uint, error)
	}

	Pokedex interface {
	}

	Service struct {
		Authorization
		Pokedex
	}
)

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
