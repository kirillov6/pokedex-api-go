package initializer

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kirillov6/pokedex-api-go/pkg/repository"
	"gorm.io/gorm"
)

func LoadEnvVariables() error {
	return godotenv.Load()
}

func ConnectToDB() (*gorm.DB, error) {
	cfg := repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	switch driver := os.Getenv("DB_DRIVERNAME"); driver {
	case "postgres":
		return repository.NewPostgresDB(cfg)
	default:
		return nil, errors.New(fmt.Sprintf("Driver %s not supported", driver))
	}
}
