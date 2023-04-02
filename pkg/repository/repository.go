package repository

import "gorm.io/gorm"

type Authorization interface {
}

type Pokedex interface {
}

type Repository struct {
	Authorization
	Pokedex
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) *Repository {
	return &Repository{DB: DB}
}
