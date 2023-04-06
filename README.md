# REST API with authentication for pokedex

## Using
* Golang ([Gin](https://gin-gonic.com) + [GORM](https://gorm.io) + [GoDotEnv](https://github.com/joho/godotenv) + [JWT](https://github.com/golang-jwt/jwt))
* Docker (postgres container)

## Run docker
`docker pull postgres`

`docker run --name=pokedex-api-db -e POSTGRES_PASSWORD='qwerty' -p 5433:5432 -d --rm postgres`

## Migrate database
`go run migrate/migrate.go`

## Run
`go run cmd/main.go`
