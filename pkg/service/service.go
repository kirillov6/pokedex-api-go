package service

import "github.com/kirillov6/pokedex-api-go/pkg/repository"

type Authorization interface {
}

type Pokedex interface {
}

type Service struct {
	Authorization
	Pokedex
	repository *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repository: repo}
}
