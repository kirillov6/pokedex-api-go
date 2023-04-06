package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kirillov6/pokedex-api-go/pkg/pokedex"
	"github.com/kirillov6/pokedex-api-go/pkg/repository"
)

const (
	tokenTTL = 12 * time.Hour
)

type (
	AuthService struct {
		repo repository.Authorization
	}

	tokenClaims struct {
		jwt.RegisteredClaims
		UserId uint `json:"user_id"`
	}
)

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user pokedex.User) (uint, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(creds pokedex.Creditionals) (string, error) {
	creds.Password = s.generatePasswordHash(creds.Password)
	user, err := s.repo.GetUser(creds)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	})

	return token.SignedString([]byte(os.Getenv("TOKEN_SALT")))
}

func (s *AuthService) ParseToken(token string) (uint, error) {
	tkn, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}

		return []byte(os.Getenv("TOKEN_SALT")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := tkn.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("Token claims has invalid type")
	}

	return claims.UserId, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_SALT"))))
}
